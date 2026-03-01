import matplotlib.pyplot as plt
import pandas as pd

# Load S3 data from CSV
df = pd.read_csv('metrics-mp.csv')

plt.style.use('dark_background')
fig, ax1 = plt.subplots(figsize=(10, 6), facecolor='#0a0a0a')
ax1.set_facecolor('#121212')

# P95 Latency (Primary Axis)
ax1.plot(df['VU'], df['p95_latency_s'], marker='o', color='#ffcc00', linewidth=3, label='p95 Latency (s)')
ax1.set_ylabel('p95 Latency (Seconds)', color='#ffcc00', fontweight='bold')
ax1.tick_params(axis='y', labelcolor='#ffcc00')
ax1.set_xlabel('Concurrent Users (VUs)')

# Throughput (Secondary Axis)
ax2 = ax1.twinx()
ax2.plot(df['VU'], df['throughput_rps'], marker='s', color='#00ff87', linewidth=3, linestyle='--', label='Throughput (RPS)')
ax2.set_ylabel('Throughput (Requests/sec)', color='#00ff87', fontweight='bold')
ax2.tick_params(axis='y', labelcolor='#00ff87')

plt.title('S3 Multipart: p95 Latency vs Throughput', fontsize=16, fontweight='bold', pad=20)
ax1.grid(True, linestyle='--', alpha=0.2)
plt.tight_layout()
plt.show()