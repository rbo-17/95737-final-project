import csv
import json
import os
import shutil
from pathlib import Path

from parse import parse

processed_results_dir = "results/processed/"
raw_results_dir = "results/raw/"
# filename = "results/raw/redis-readheavy-img-239-1721347791.csv"

results_files = os.listdir(raw_results_dir)


def pretty_print(input: dict | list[dict]):
    print(json.dumps(input, sort_keys=True, indent=4))


dirpath = Path(processed_results_dir)
if dirpath.exists() and dirpath.is_dir():
    shutil.rmtree(dirpath)

dirpath.mkdir(parents=True, exist_ok=True)


# class RequestData:
#     def __init__(self):
#
#
# class DbSummary:
#
#     def __init__(self):
#         self.read_heavy_metadata = dict()
#         self.balanced_metadata = dict()
#         self.write_heavy_metadata = dict()
#
#     def append_results(self, test_type: str, meta_data: dict):
#         if test_type == "readheavy":
#             self.read_heavy_metadata = meta_data
#         elif test_type == "balanced":
#             self.balanced_metadata = meta_data
#         elif test_type == "writeheavy":
#             self.write_heavy_metadata = meta_data
#         else:
#             raise Exception("Test type not recognized!")
#
#     def get_read_heavy(self) -> dict:
#         return self.read_heavy_metadata
#
#     def get_balanced(self):
#         return self.balanced_metadata
#
#     def get_write_heavy(self):
#         return self.write_heavy_metadata
#
#     def get_latency(self):
#         pass
#
#
# db_summaries = {
#     "cassandra": DbSummary(),
#     "mongodb": DbSummary(),
#     "mysql": DbSummary(),
#     "redis": DbSummary(),
# }

db_summaries = {
    "readheavy": list(),
    "balanced": list(),
    "writeheavy": list()
}

# db_summaries = dict()

for results_file in results_files:

    results = list()

    with open(raw_results_dir + results_file, mode="r") as file:
        csv_file = csv.DictReader(file)
        for line in csv_file:
            results.append(line)

    error_count = 0
    for result in results:
        if result["Error"] != "nil" or result["Ok"] != "true":
            error_count += 1

    # error_count = reduce(lambda s, r: s if r["Error"] == "nil" and r["Ok"] == "true" else s + 1, results)

    latencies = list(map(lambda r: int(r["Latency"]), results))
    avg_latency = sum(latencies) / len(latencies)

    parse_res = parse('{}-{}-{}-{}-{}.csv', results_file)
    test_metadata = {
        "filename": results_file,
        "db_name": parse_res[0],
        "result_name": parse_res[0] + "/" + parse_res[2],
        "test_type": parse_res[1],
        "data_type": parse_res[2],
        "duration_ms": int(parse_res[3]),
        "exec_time": parse_res[4],
        "op_count": len(results),
        "error_count": error_count,
        "avg_latency": avg_latency
    }

    ops_per_ms = test_metadata["op_count"] / int(test_metadata["duration_ms"])
    test_metadata["ops_per_second"] = int(ops_per_ms * 1000.0)

    # if int(test_metadata["duration_ms"]) > 0:
    #     print(int(test_metadata["duration_ms"]))
    #     duration_s = int(test_metadata["duration_ms"] / 1000)
    #     test_metadata["ops_per_second"] = test_metadata["op_count"] / duration_s
    #
    # else:
    #     test_metadata["ops_per_second"] = 0

    db_summaries[test_metadata["test_type"]].append(test_metadata)
    # db_summaries[test_metadata["test_type"]].append_results(test_metadata)


def write_csv(fname: str, keys: list[str], data: list[dict]):
    with open(processed_results_dir + fname + '.csv', 'w', newline='') as csvfile:
        # print("data")
        # pretty_print(data)
        fieldnames = keys
        writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(data)


# print("db_summaries")
# pretty_print(db_summaries)

# Write all summaries to file
for key, summaries in db_summaries.items():
    sorted_summaries = sorted(summaries, key=lambda d: d['result_name'], reverse=True)
    write_csv(key, summaries[0].keys(), sorted_summaries)

# Get latencies and write to file
latency_summaries = list()
for key, summaries in db_summaries.items():
    latency_summaries += summaries
# print(type(latency_summaries[0][0]))
# pretty_print(latency_summaries[0][0])
# print(latency_summaries[0][0].keys())
write_csv("latencies", latency_summaries[0].keys(), latency_summaries)
