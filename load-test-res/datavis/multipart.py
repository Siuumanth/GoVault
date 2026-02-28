import matplotlib.pyplot as plt
import seaborn as sns
import pandas as pd

# Data from S3 Multipart Tests (1-6)
s3_data = {
    'VUs': [200, 400, 600, 750, 1000],
    'Throughput': [174.03, 202.29, 193.75, 157.69, 162.80],
    'p95_Latency': [1.33, 4.15, 6.02, 10.00, 13.32]
}
df_s3 = pd.DataFrame(s3_data)

# Dark Mode Styling
plt.style.use('dark_background')
fig, ax1 = plt.subplots(figsize=(10, 6), facecolor='#121212')
ax1.set_facecolor('#1e1e1e')
ax1.grid(True, linestyle='--', alpha=0.5)


# Primary Axis: Throughput (Neon Green)
tp_color = '#00ff87'
ln1 = ax1.plot(df_s3['VUs'], df_s3['Throughput'], color=tp_color, marker='o', 
               markersize=10, linewidth=3, label='Throughput (Req/s)')
ax1.set_ylabel('Requests per Second', color=tp_color, fontsize=12, fontweight='bold')
ax1.tick_params(axis='y', labelcolor=tp_color)
ax1.set_xlabel('Concurrent Users (VUs)', fontsize=12, color='white')

# Secondary Axis: Latency (Amber)
ax2 = ax1.twinx()
lat_color = '#ffcc00'
ln2 = ax2.plot(df_s3['VUs'], df_s3['p95_Latency'], color=lat_color, marker='s', 
               markersize=10, linewidth=3, linestyle='--', label='p95 Latency (s)')
ax2.set_ylabel('p95 Latency (Seconds)', color=lat_color, fontsize=12, fontweight='bold')
ax2.tick_params(axis='y', labelcolor=lat_color)

# Title & Legend
plt.title('S3 Multipart Performance: High Scale Efficiency', fontsize=16, pad=20, fontweight='bold')
lines = ln1 + ln2
labels = [l.get_label() for l in lines]
ax1.legend(lines, labels, loc='upper left', frameon=True, facecolor='#1e1e1e', edgecolor='#333333')

plt.tight_layout()
plt.show()