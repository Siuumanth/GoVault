import matplotlib.pyplot as plt
import pandas as pd

# Load data
df = pd.read_csv('metrics-proxy.csv')

plt.style.use('dark_background')
fig, ax1 = plt.subplots(figsize=(12, 7), facecolor='#0a0a0a')
ax1.set_facecolor('#121212')

# Throughput - Increased linewidth and marker size
ax1.plot(df['VU'], df['throughput_rps'], marker='o', color='#00ff87', 
         linewidth=4, markersize=8, label='Throughput (RPS)', zorder=3)
ax1.set_xlabel('Concurrent Virtual Users (VUs)', fontsize=14, fontweight='bold', labelpad=15)
ax1.set_ylabel('Throughput (Requests/sec)', color='#00ff87', fontsize=14, fontweight='bold')
ax1.tick_params(axis='both', labelcolor='#ffffff', labelsize=12)

# Error Rate - Secondary Axis
ax2 = ax1.twinx()
ax2.plot(df['VU'], df['error_rate_percent'], marker='s', color='#ff4444', 
         linewidth=4, markersize=8, linestyle='--', label='Error Rate (%)', zorder=3)
ax2.set_ylabel('Error Rate (%)', color='#ff4444', fontsize=14, fontweight='bold')
ax2.tick_params(axis='y', labelcolor='#ff4444', labelsize=12)

# Title and Grid
plt.title('Proxy: Throughput vs Error Rate', fontsize=18, fontweight='bold', pad=25, color='#ffffff')
ax1.grid(True, linestyle=':', alpha=0.3, zorder=1)

plt.tight_layout()
plt.show()