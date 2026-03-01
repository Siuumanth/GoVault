Contents:

introduction (what is this, tools used, context, overview of architecture)
and sayying strictly this test is based on locally tested , on windows ,  ryzen 5, 16 gb ram so there is wsl overhead also , so is less efifcient than actual cloud 
things to know before hand , virtual users vs actual users, mapping between them, what is p95 n stuff, why errors can happen in http  , like tcp connections , db pool exhaust n stuff 


# **1. Introduction**

### **Project Overview**

This report covers the performance testing of **GoVault**, a cloud-native file storage system built with **Go (Golang)**. The tests compare two ways of uploading files: **Proxy-Based-Chunked** (through the backend) and **Direct-to-S3** (using MinIO).
### **System Architecture**

The system uses a **microservices** design. To test this, we simulated a full production stack using **Docker Compose**, running multiple containers simultaneously:
- **4 Go Services:** Gateway, Auth, Upload, and Files.
- **3 PostgreSQL Instances:** Dedicated database for each core service.
- **Object Storage:** **MinIO** (S3-compatible) for file storage.
- **Observability:** **Prometheus** and **Grafana** for tracking metrics.
### **Tools Used**
- **k6:** For simulating concurrent users and file uploads.
- **Prometheus & Grafana:** For monitoring system health and speed.
- **Zap:** For high-performance structured logging.

### **Testing Environment & Limits**

> **Strictly Local Test:** These tests were run on a single machine (**AMD Ryzen 5, 16GB RAM**) using **Windows with WSL2**.
> 
> Running **10+ containers** (4 services, 3 databases, MinIO, Prometheus, and Grafana) on one machine creates significant CPU and I/O overhead due to the WSL2 virtualization layer. This makes the system less efficient than a real cloud environment, but it accurately shows how the architecture handles stress


---


