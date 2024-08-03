import csv
import functools
import json
import os
import shutil
from pathlib import Path
from typing import Any

from parse import parse

processed_results_dir = "results/processed/"
raw_results_dir = "results/raw/"

results_files = os.listdir(raw_results_dir)


def pretty_print(input: dict | list[dict]):
    print(json.dumps(input, sort_keys=True, indent=4))


dirpath = Path(processed_results_dir)
if dirpath.exists() and dirpath.is_dir():
    shutil.rmtree(dirpath)

dirpath.mkdir(parents=True, exist_ok=True)

DATA_TYPE_SM_TXT = "sm"
DATA_TYPE_LG_TXT = "lg"
DATA_TYPE_IMG = "img"

DN_MAP = dict()
DN_MAP[DATA_TYPE_SM_TXT] = {
    "min": 100,
    "max": 200
}
DN_MAP[DATA_TYPE_LG_TXT] = {
    "min": 10000,
    "max": 20000
}
DN_MAP[DATA_TYPE_IMG] = {
    "min": 1000000,
    "max": 2000000
}


def calculate_denormalization_entries(data_type: str, results: list[dict[str, str]]) -> list[dict[str, Any]]:
    sorted_results = dict()
    data_type_mapping = DN_MAP[data_type]
    for result in results:
        df = int(int(result["Size(B)"]) / data_type_mapping["max"]) + 1

        if df not in sorted_results:
            sorted_results[df] = list()

        sorted_results[df].append(result)

    denormalization_summaries = list()
    for df, results in sorted_results.items():

        total_ops = len(results)
        total_bytes = 0
        total_latency_us = 0

        for result in results:
            total_bytes += int(result["Size(B)"])
            total_latency_us += int(result["Latency"])

        ops_per_us = total_ops / total_latency_us
        ops_per_s = int(ops_per_us * 1000000)

        bytes_per_us = total_bytes / total_latency_us
        bytes_per_s = int(bytes_per_us * 1000000)

        denormalization_summaries.append({
            "data_type": data_type,
            "denormalization_factor": df,
            "total_ops": total_ops,
            "ops_per_second": ops_per_s,
            "total_bytes": total_bytes,
            "bytes_per_second": bytes_per_s,
            "total_latency(µs)": total_latency_us,
            "total_latency(ms)": int(total_latency_us * 1000)
        })

    return denormalization_summaries


db_summaries = {
    "readheavy": list(),
    "balanced": list(),
    "writeheavy": list()
}

denormalization_summaries = None
for results_file in results_files:

    results = list()

    # Read all results (entries) in file
    with open(raw_results_dir + results_file, mode="r") as file:
        csv_file = csv.DictReader(file)
        for line in csv_file:
            results.append(line)

    # Sum error count
    error_count = 0
    for result in results:
        if result["Error"] != "nil" or result["Ok"] != "true":
            error_count += 1

    # Calculate average latency
    latencies = list(map(lambda r: int(r["Latency"]), results))
    avg_latency = sum(latencies) / len(latencies)

    # Build summary
    # Check if opts are provided
    try:
        parse_res = parse('{}-{}-{}-{}-{}-{}.csv', results_file)
        denormalization_factor = int(parse('df{}', parse_res[5])[0])  # Expand as necessary via a parse_opts() function

    except TypeError:
        parse_res = parse('{}-{}-{}-{}-{}.csv', results_file)
        denormalization_factor = 1

    test_summary = {
        "filename": results_file,
        "db_name": parse_res[0],
        "result_name": parse_res[0] + "/" + parse_res[2],
        "test_type": parse_res[1],
        "data_type": parse_res[2],
        "duration_ms": int(parse_res[3]),
        "exec_time": parse_res[4],
        "op_count": len(results),
        "error_count": error_count,
        "avg_latency": avg_latency,
        "denormalization_factor": denormalization_factor
    }

    # Calculate operations per second (ops/s)
    ops_per_ms = test_summary["op_count"] / int(test_summary["duration_ms"])
    test_summary["ops_per_second"] = int(ops_per_ms * 1000.0)

    # Calculate bytes per second (B/s)
    bytes_processed = functools.reduce(lambda s, r: s + int(r["Size(B)"]), results, 0)
    bytes_per_ms = bytes_processed / int(test_summary["duration_ms"])
    test_summary["bytes_per_second"] = int(bytes_per_ms * 1000.0)
    test_summary["mb_per_second"] = round(test_summary["bytes_per_second"] / 2 ** 20, 2)

    # Assign a denormalization key to make sorting easier
    denormalization_key_mapping = {
        "cassandra": 100,
        "mongodb": 200,
        "mysql": 300,
        "redis": 400
    }
    test_summary["denormalization_key"] = (
            denormalization_key_mapping[test_summary["db_name"]] +
            test_summary["denormalization_factor"])

    # Append result to summaries list
    db_summaries[test_summary["test_type"]].append(test_summary)

    # # Calculate denormalization entries
    # denormalization_summaries = calculate_denormalization_entries(test_summary["data_type"], results)


def write_csv(fname: str, keys: list[str], data: list[dict]):
    with open(processed_results_dir + fname + '.csv', 'w', newline='') as csvfile:
        fieldnames = keys
        writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(data)


# Write all summaries to file
for key, summaries in db_summaries.items():
    sorted_summaries = sorted(summaries, key=lambda d: d['result_name'], reverse=True)
    write_csv(key, summaries[0].keys(), sorted_summaries)

    sorted_denormalization_summaries = sorted(summaries, key=lambda d: d['denormalization_key'])
    write_csv("denorm" + key, summaries[0].keys(), sorted_denormalization_summaries)

# Get latencies and write to file
latency_summaries = list()
for key, summaries in db_summaries.items():
    latency_summaries += summaries

write_csv("latencies", latency_summaries[0].keys(), latency_summaries)
