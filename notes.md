# Flags

Set a prefix for logs while initing - If users are combining logs they can distinguish better.

# Fallbacks and Indirect Signals
1. TODO: Fallback connections - Allows working by getting indirect signals from computes instead of server.
Also will update the Network.Connect field if needed.

2. REV: Assuming leadership - If server is unreachable for many nodes, allow them to setup their own network. Will require knowing history of past connections to computes.

# Advanced logging
Chapter 13 of Network programming with Go covers this, for performance and controlling verbosity.
Add metrics, leveled logs, on demand debug logs, anomaly reporting.