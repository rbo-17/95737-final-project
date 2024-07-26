import csv
import os

import matplotlib.pyplot as plt
import pandas as pd
import seaborn as sns
from parse import parse

raw_results_dir = "results/raw/"
# filename = "results/raw/redis-readheavy-img-239-1721347791.csv"

results_files = os.listdir(raw_results_dir)

results = list()
for results_file in results_files:

    parse_res = parse('{}-{}-{}-{}-{}.csv', results_file)
    test_metadata = {
        "db_name": parse_res[0],
        "test_type": parse_res[1],
        "data_type": parse_res[2],
        "duration_ms": parse_res[3],
        "exec_time": parse_res[4]
    }

    with open(raw_results_dir + results_file, mode="r") as file:
        csvFile = csv.DictReader(file)
        for line in csvFile:
            # line["Test type"] = test_metadata["test_type"]
            line["Test type"] = test_metadata["test_type"] + "/" + test_metadata["data_type"]
            line["db_name"] = test_metadata["db_name"]
            results.append(line)
            # print(line)

sns.set_theme(style="whitegrid", palette="tab10")
# sns.set_theme(style="ticks", palette="tab10")


# sns.countplot(x="Column", data=ds)

results_df = pd.DataFrame(results)

plt.figure(figsize=(16,12)) # this creates a figure 8 inch wide, 4 inch high
bp = sns.boxplot(
    x="Test type",
    y="Latency",
    hue="db_name",
    palette=["red", "blue", "cyan", "green"],
    data=results_df)
# sns.despine(offset=10, trim=True)
bp.set_xticklabels(bp.get_xticklabels(), rotation=30)

# plt.show()
plt.savefig("out.png")
