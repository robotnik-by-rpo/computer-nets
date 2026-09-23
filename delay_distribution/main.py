import matplotlib.pyplot as plt
import numpy as np

times = []
with open("ping.txt", "r", encoding="utf-8") as pings:
    for p in pings:
        if "time=" in p:
            part = p.split("=")[-1].split(" ")[0]
            times.append(part)

times = np.array(times,dtype=int)
plt.hist(times, bins=int(np.sqrt(1000)),density=True)
plt.show()