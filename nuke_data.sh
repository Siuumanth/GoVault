#!/bin/bash

# --- Database Config ---
# We use the service names from your docker-compose.yml
# Change 'auth-db' etc. to match your actual service names if different.

echo "🚀 Starting GoVault Data Nuke..."

# 1. AUTH DB
echo "🧹 Cleaning Auth DB..."
docker exec -e PGPASSWORD=authpass -it auth-db psql -U auth -d auth_db -c "TRUNCATE TABLE users CASCADE;"

# 2. UPLOAD DB
echo "🧹 Cleaning Upload DB..."
# Truncating upload_sessions handles upload_chunks automatically due to CASCADE
docker exec -e PGPASSWORD=uploadpass -it upload-db psql -U upload -d upload_db -c "TRUNCATE TABLE upload_sessions CASCADE;"

# 3. FILES DB
echo "🧹 Cleaning Files DB..."
# Truncating files handles file_shares, file_shortcuts, and public_files automatically
docker exec -e PGPASSWORD=filespass -it files-db psql -U files -d files_db -c "TRUNCATE TABLE files CASCADE;"

echo "✨ All databases cleared. Your GoVault is fresh for testing."
