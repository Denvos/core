package headers

import (
    "fmt"
    "runtime"
)

func BuildUserAgent(app, version, os, arch string) string {
    if os == "" {
        os = runtime.GOOS
    }
    if arch == "" {
        arch = runtime.GOARCH
    }
    if app == "" {
        app = "Denvos"
    }
    if version == "" {
        version = "0.1"
    }
    return fmt.Sprintf("%s/%s (%s; %s)", app, version, os, arch)
}

func ParseUserAgent(ua string) (app, version, os, arch string) {
    // Very basic parsing: "Denvos/0.1 (linux; amd64)"
    if ua == "" {
        return
    }
    // Extract app/version
    idx := strings.Index(ua, " ")
    if idx == -1 {
        return
    }
    appVer := ua[:idx]
    parts := strings.SplitN(appVer, "/", 2)
    if len(parts) == 2 {
        app = parts[0]
        version = parts[1]
    }
    // Extract OS/Arch from parentheses
    start := strings.Index(ua, "(")
    end := strings.Index(ua, ")")
    if start != -1 && end != -1 && end > start {
        info := ua[start+1 : end]
        infos := strings.Split(info, ";")
        if len(infos) >= 2 {
            os = strings.TrimSpace(infos[0])
            arch = strings.TrimSpace(infos[1])
        }
    }
    return
}
