import matplotlib.pyplot as plt
import pandas as pd

# Load data from CSV
df = pd.read_csv('metrics-proxy.csv')

plt.style.use('dark_background')
fig, ax1 = plt.subplots(figsize=(10, 6), facecolor='#0a0a0a')
ax1.set_facecolor('#121212')

# Throughput (Primary Axis)
ax1.plot(df['VU'], df['throughput_rps'], marker='o', color='#00ff87', linewidth=3, label='Throughput (RPS)')
ax1.set_ylabel('Throughput (Requests/sec)', color='#00ff87', fontweight='bold')
ax1.tick_params(axis='y', labelcolor='#00ff87')
ax1.set_xlabel('Concurrent Users (VUs)')

# Error Rate (Secondary Axis)
ax2 = ax1.twinx()
ax2.plot(df['VU'], df['error_rate_percent'], marker='s', color='#ff4444', linewidth=3, linestyle='--', label='Error Rate (%)')
ax2.set_ylabel('Error Rate (%)', color='#ff4444', fontweight='bold')
ax2.tick_params(axis='y', labelcolor='#ff4444')

plt.title('Proxy: Throughput vs Error Rate', fontsize=16, fontweight='bold', pad=20)
ax1.grid(True, linestyle='--', alpha=0.2) # Subtle grid
plt.tight_layout()
plt.show()