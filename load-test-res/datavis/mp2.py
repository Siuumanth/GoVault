import matplotlib.pyplot as plt
import pandas as pd

df = pd.read_csv('metrics-mp.csv')

plt.style.use('dark_background')
fig, ax1 = plt.subplots(figsize=(12, 7), facecolor='#0a0a0a')
ax1.set_facecolor('#121212')

# P95 Latency
ax1.plot(df['VU'], df['p95_latency_s'], marker='o', color='#ffcc00', 
         linewidth=4, markersize=8, label='p95 Latency (s)', zorder=3)
ax1.set_xlabel('Concurrent Virtual Users (VUs)', fontsize=14, fontweight='bold', labelpad=15)
ax1.set_ylabel('p95 Latency (Seconds)', color='#ffcc00', fontsize=14, fontweight='bold')
ax1.tick_params(axis='both', labelcolor='#ffffff', labelsize=12)

# Throughput
ax2 = ax1.twinx()
ax2.plot(df['VU'], df['throughput_rps'], marker='s', color='#00ff87', 
         linewidth=4, markersize=8, linestyle='--', label='Throughput (RPS)', zorder=3)
ax2.set_ylabel('Throughput (Requests/sec)', color='#00ff87', fontsize=14, fontweight='bold')
ax2.tick_params(axis='y', labelcolor='#00ff87', labelsize=12)

plt.title('S3 Multipart: p95 Latency vs Throughput', fontsize=18, fontweight='bold', pad=25, color='#ffffff')
ax1.grid(True, linestyle=':', alpha=0.3, zorder=1)

plt.tight_layout()
plt.show()