import matplotlib.pyplot as plt
import pandas as pd

df_mp = pd.read_csv('metrics-mp.csv')
df_proxy = pd.read_csv('metrics-proxy.csv')

plt.style.use('dark_background')
fig, ax1 = plt.subplots(figsize=(14, 8), facecolor='#0a0a0a')
ax1.set_facecolor('#121212')

# Left Axis: Throughput (Solid Lines)
ax1.plot(df_mp['VU'], df_mp['throughput_rps'], marker='o', color='#00ff87', linewidth=4, label='S3: Throughput (RPS)')
ax1.plot(df_proxy['VU'], df_proxy['throughput_rps'], marker='o', color='#00d4ff', linewidth=4, label='Proxy: Throughput (RPS)')
ax1.set_ylabel('Throughput (Requests/sec)', color='#ffffff', fontsize=14, fontweight='bold')
ax1.set_xlabel('Concurrent Virtual Users (VUs)', fontsize=14, fontweight='bold', color='#ffffff')

# Right Axis: P95 Latency (Dashed Lines)
ax2 = ax1.twinx()
ax2.plot(df_mp['VU'], df_mp['p95_latency_s'], marker='s', color='#ffcc00', linewidth=3, linestyle='--', label='S3: p95 Latency (s)')
ax2.plot(df_proxy['VU'], df_proxy['p95_latency_s'], marker='s', color='#ff4444', linewidth=3, linestyle='--', label='Proxy: p95 Latency (s)')
ax2.set_ylabel('p95 Latency (Seconds)', color='#ffffff', fontsize=14, fontweight='bold')

plt.title('Performance Showdown: Throughput & Latency', fontsize=20, fontweight='bold', pad=30)
ax1.grid(True, linestyle=':', alpha=0.2)
# Merge legends
lines1, labels1 = ax1.get_legend_handles_labels()
lines2, labels2 = ax2.get_legend_handles_labels()
ax1.legend(lines1 + lines2, labels1 + labels2, loc='upper left', framealpha=0.3)

plt.tight_layout()
plt.show()