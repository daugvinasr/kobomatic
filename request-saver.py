from mitmproxy import ctx
import mitmproxy.http
import os
from datetime import datetime

# mitmproxy -s request-saver.py --mode reverse:http://192.168.1.55:8085 --listen-host 192.168.1.46 --listen-port 8084
# mitmproxy -s request-saver.py --mode reverse:https://storeapi.kobo.com --listen-host 192.168.1.46 --listen-port 8084

class Logger:
    def __init__(self):
        self.log_dir = "logs"
        os.makedirs(self.log_dir, exist_ok=True)
        self.flows = {}

    def request(self, flow: mitmproxy.http.HTTPFlow):
        self.flows[flow] = {
            "timestamp": datetime.now().strftime("%Y%m%d_%H%M%S_%f"),
            "host": flow.request.host,
            "request": self.format_request(flow)
        }

    def response(self, flow: mitmproxy.http.HTTPFlow):
        if flow in self.flows:
            self.flows[flow]["response"] = self.format_response(flow)
            self.log_flow(flow)
            del self.flows[flow]

    def format_request(self, flow: mitmproxy.http.HTTPFlow) -> str:
        output = f"REQUEST: {flow.request.method} {flow.request.url}\n"
        output += "Headers:\n"
        for key, value in flow.request.headers.items():
            output += f"{key}: {value}\n"
        output += "\nBody:\n"
        output += flow.request.content.decode(errors="replace")
        return output

    def format_response(self, flow: mitmproxy.http.HTTPFlow) -> str:
        output = f"RESPONSE: {flow.response.status_code}\n"
        output += "Headers:\n"
        for key, value in flow.response.headers.items():
            output += f"{key}: {value}\n"
        output += "\nBody:\n"
        output += flow.response.content.decode(errors="replace")
        return output

    def log_flow(self, flow: mitmproxy.http.HTTPFlow):
        flow_data = self.flows[flow]
        filename = f"{flow_data['timestamp']}_{flow_data['host']}.txt"
        filepath = os.path.join(self.log_dir, filename)

        with open(filepath, "w", encoding="utf-8") as f:
            f.write(flow_data["request"])
            f.write("\n\n" + "="*50 + "\n\n")
            f.write(flow_data["response"])

addons = [Logger()]