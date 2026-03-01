import matplotlib.pyplot as plt
import pandas as pd

df = pd.read_csv('metrics-proxy.csv')

plt.style.use('dark_background')
fig, ax1 = plt.subplots(figsize=(12, 7), facecolor='#0a0a0a')
ax1.set_facecolor('#121212')

# P95 Latency (Yellow) - Note: converting median_ms to seconds for comparison
ax1.plot(df['VU'], df['p95_latency_s'], marker='o', color='#ffcc00', 
         linewidth=4, markersize=8, label='P95 Latency (Tail)', zorder=3)

# Median Latency (Cyan) - Converting ms to s
ax1.plot(df['VU'], df['median_latency_ms'] / 1000, marker='D', color='#00d4ff', 
         linewidth=3, linestyle='-.', markersize=7, label='Median Latency (Average)', zorder=3)

ax1.set_xlabel('Concurrent Virtual Users (VUs)', fontsize=14, fontweight='bold', labelpad=15, color='#ffffff')
ax1.set_ylabel('Latency (Seconds)', color='#ffffff', fontsize=14, fontweight='bold')
ax1.tick_params(axis='both', labelcolor='#ffffff', labelsize=12)

plt.title('Latency Spread: Median vs P95 - Proxy', fontsize=18, fontweight='bold', pad=25, color='#ffffff')
ax1.grid(True, linestyle=':', alpha=0.3, zorder=1)
ax1.legend(fontsize=12, loc='upper left', framealpha=0.2)

plt.tight_layout()
plt.show()