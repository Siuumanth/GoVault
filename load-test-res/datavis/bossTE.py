import matplotlib.pyplot as plt
import pandas as pd

# Load the datasets
df_mp = pd.read_csv('metrics-mp.csv')
df_proxy = pd.read_csv('metrics-proxy.csv')

plt.style.use('dark_background')
fig, ax1 = plt.subplots(figsize=(14, 8), facecolor='#0a0a0a')
ax1.set_facecolor('#121212')

# --- LEFT AXIS: THROUGHPUT (RPS) ---
# S3 Throughput (Neon Green)
ax1.plot(df_mp['VU'], df_mp['throughput_rps'], marker='o', color='#00ff87', 
         linewidth=4, markersize=10, label='S3: Throughput (RPS)', zorder=3)
# Proxy Throughput (Cyan)
ax1.plot(df_proxy['VU'], df_proxy['throughput_rps'], marker='o', color='#00d4ff', 
         linewidth=4, markersize=10, label='Proxy: Throughput (RPS)', zorder=3)

ax1.set_xlabel('Concurrent Virtual Users (VUs)', fontsize=14, fontweight='bold', labelpad=15, color='#ffffff')
ax1.set_ylabel('Throughput (Requests/sec)', color='#ffffff', fontsize=14, fontweight='bold')
ax1.tick_params(axis='both', labelcolor='#ffffff', labelsize=12)

# --- RIGHT AXIS: ERROR RATE (%) ---
ax2 = ax1.twinx()
# S3 Error Rate (Bright Red Dashed)
ax2.plot(df_mp['VU'], df_mp['error_rate_percent'], marker='x', color='#ff4444', 
         linewidth=3, linestyle='--', markersize=10, label='S3: Error Rate (%)', zorder=3)
# Proxy Error Rate (Orange Dashed)
ax2.plot(df_proxy['VU'], df_proxy['error_rate_percent'], marker='x', color='#ff8c00', 
         linewidth=3, linestyle='--', markersize=10, label='Proxy: Error Rate (%)', zorder=3)

ax2.set_ylabel('Error Rate (%)', color='#ffffff', fontsize=14, fontweight='bold')
ax2.tick_params(axis='y', labelcolor='#ffffff', labelsize=12)

# Title and Grid
plt.title('Reliability Benchmark: Throughput vs Error Rate', fontsize=20, fontweight='bold', pad=30, color='#ffffff')
ax1.grid(True, linestyle=':', alpha=0.2, zorder=1)

# Combined Legend for both axes
lines1, labels1 = ax1.get_legend_handles_labels()
lines2, labels2 = ax2.get_legend_handles_labels()
ax1.legend(lines1 + lines2, labels1 + labels2, loc='upper left', fontsize=11, framealpha=0.2)

plt.tight_layout()
plt.show()