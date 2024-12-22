# Install

```bash
wget https://github.com/javilobo8/temperature-exporter/releases/download/v0.0.1/temperature-exporter -O /tmp/temperature-exporter
chmod +x /tmp/temperature-exporter
mv /tmp/temperature-exporter /usr/local/bin/
```

# Create service

```bash

cat <<EOF > /etc/systemd/system/temperature-exporter.service
[Unit]
Description=Temperature Exporter Service
After=network.target

[Service]
ExecStart=/usr/local/bin/temperature-exporter
Restart=always
User=nobody
Group=nogroup

[Install]
WantedBy=multi-user.target

EOF

systemctl daemon-reload
systemctl enable temperature-exporter
systemctl start temperature-exporter
```