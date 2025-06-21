#!/bin/sh

echo "⏳ Menunggu PostgreSQL di $DB_HOST:$DB_PORT..."

while ! nc -z "$DB_HOST" "$DB_PORT"; do
  sleep 1
done

echo "✅ Database sudah siap, menjalankan aplikasi..."
exec "$@"
