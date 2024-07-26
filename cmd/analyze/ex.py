"""
Grouped barplots
================

_thumb: .36, .5
"""
import csv

import matplotlib.pyplot as plt
import seaborn as sns

sns.set_theme(style="ticks", palette="pastel")

# Load the example tips dataset
tips = sns.load_dataset("tips")

# Draw a nested boxplot to show bills by day and time
bp = sns.boxplot(
    x="day",
    y="total_bill",
    hue="smoker",
    palette=["m", "g"],
    data=tips)
# sns.despine(offset=10, trim=True)

plt.savefig("out.png")
