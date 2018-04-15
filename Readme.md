(types.Container) {
 ID: (string) (len=64) "87b1485a13d2c20ff417ebecc300162046e5198ff2bbe247f6be7cfedb414706",
 Names: ([]string) (len=1 cap=4) {
  (string) (len=23) "/clubaccounting_redis_1"
 },
 Image: (string) (len=18) "redis:3.2.7-alpine",
 ImageID: (string) (len=71) "sha256:147b1b8460d23e3d8a24b3503a9947ffdee6aaf520e4d41cfd8b045715809188",
 Command: (string) (len=48) "docker-entrypoint.sh redis-server --bind 0.0.0.0",
 Created: (int64) 1512864167,
 Ports: ([]types.Port) (len=1 cap=4) {
  (types.Port) {
   IP: (string) (len=7) "0.0.0.0",
   PrivatePort: (uint16) 6379,
   PublicPort: (uint16) 6380,
   Type: (string) (len=3) "tcp"
  }
 },
 SizeRw: (int64) 0,
 SizeRootFs: (int64) 0,
 Labels: (map[string]string) (len=6) {
  (string) (len=35) "com.docker.compose.container-number": (string) (len=1) "1",
  (string) (len=25) "com.docker.compose.oneoff": (string) (len=5) "False",
  (string) (len=26) "com.docker.compose.project": (string) (len=14) "clubaccounting",
  (string) (len=26) "com.docker.compose.service": (string) (len=5) "redis",
  (string) (len=26) "com.docker.compose.version": (string) (len=6) "1.10.0",
  (string) (len=30) "com.docker.compose.config-hash": (string) (len=64) "8f764167a730410396da6ecb0518d10b1d9f1faad47cf9bb4cdd643f8c2f2d2f"
 },
 State: (string) (len=7) "running",
 Status: (string) (len=10) "Up 2 hours",
 HostConfig: (struct { NetworkMode string "json:\",omitempty\"" }) {
  NetworkMode: (string) (len=22) "clubaccounting_default"
 },
 NetworkSettings: (*types.SummaryNetworkSettings)(0xc420324048)({
  Networks: (map[string]*network.EndpointSettings) (len=1) {
   (string) (len=22) "clubaccounting_default": (*network.EndpointSettings)(0xc42033a240)({
    IPAMConfig: (*network.EndpointIPAMConfig)(<nil>),
    Links: ([]string) <nil>,
    Aliases: ([]string) <nil>,
    NetworkID: (string) (len=64) "21b5a7a9d682d18ce6ffab8cc04ce07b6fc58953c41070c64d74a54dc0fc8f06",
    EndpointID: (string) (len=64) "2e9063c2632e1d1dce89914ffb67b6466eabc3e551a5a1536443fd6cfc5b3676",
    Gateway: (string) (len=10) "172.20.0.1",
    IPAddress: (string) (len=10) "172.20.0.4",
    IPPrefixLen: (int) 16,
    IPv6Gateway: (string) "",
    GlobalIPv6Address: (string) "",
    GlobalIPv6PrefixLen: (int) 0,
    MacAddress: (string) (len=17) "02:42:ac:14:00:04"
   })
  }
 }),
 Mounts: ([]types.MountPoint) (len=1 cap=4) {
  (types.MountPoint) {
   Type: (mount.Type) (len=4) "bind",
   Name: (string) "",
   Source: (string) (len=16) "/opt/xredis/data",
   Destination: (string) (len=5) "/data",
   Driver: (string) "",
   Mode: (string) (len=2) "rw",
   RW: (bool) true,
   Propagation: (mount.Propagation) ""
  }
 }
}
