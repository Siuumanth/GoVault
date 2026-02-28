import matplotlib.pyplot as plt
import seaborn as sns
import pandas as pd

# Data from Proxy Tests (1-5)
proxy_data = {
    'VUs': [200, 300, 400],
    'Throughput': [91.67, 134.85, 140.43],
    'p95_Latency': [3.85, 3.39, 5.50]
}
df_p = pd.DataFrame(proxy_data)

# Dark Mode Styling
plt.style.use('dark_background')
fig, ax1 = plt.subplots(figsize=(10, 6), facecolor='#121212')
ax1.set_facecolor('#1e1e1e')
ax1.grid(True, linestyle='--', alpha=0.5)

# Primary Axis: Throughput (Cyan)
tp_color = '#00d4ff'
ln1 = ax1.plot(df_p['VUs'], df_p['Throughput'], color=tp_color, marker='o', 
               markersize=10, linewidth=3, label='Throughput (Req/s)')
ax1.set_ylabel('Requests per Second', color=tp_color, fontsize=12, fontweight='bold')
ax1.tick_params(axis='y', labelcolor=tp_color)
ax1.set_xlabel('Concurrent Users (VUs)', fontsize=12, color='white')

# Secondary Axis: Latency (Soft Red)
ax2 = ax1.twinx()
lat_color = '#ff4b5c'
ln2 = ax2.plot(df_p['VUs'], df_p['p95_Latency'], color=lat_color, marker='s', 
               markersize=10, linewidth=3, linestyle='--', label='p95 Latency (s)')
ax2.set_ylabel('p95 Latency (Seconds)', color=lat_color, fontsize=12, fontweight='bold')
ax2.tick_params(axis='y', labelcolor=lat_color)

# Title & Legend
plt.title('Proxy Upload Performance: Throughput vs Latency', fontsize=16, pad=20, fontweight='bold')
lines = ln1 + ln2
labels = [l.get_label() for l in lines]
ax1.legend(lines, labels, loc='upper left', frameon=True, facecolor='#1e1e1e', edgecolor='#333333')

plt.tight_layout()
plt.show()
