package headers

import (
    "strconv"
    "strings"
)

type AcceptItem struct {
    Type    string
    Subtype string
    Quality float64
    Params  map[string]string
}

func (a AcceptItem) String() string {
    s := a.Type + "/" + a.Subtype
    if a.Quality > 0 && a.Quality < 1 {
        s += ";q=" + strconv.FormatFloat(a.Quality, 'f', -1, 64)
    }
    for k, v := range a.Params {
        if k != "q" {
            s += ";" + k + "=" + v
        }
    }
    return s
}

func ParseAccept(header string) []AcceptItem {
    if header == "" {
        return nil
    }
    var items []AcceptItem
    for _, part := range strings.Split(header, ",") {
        part = strings.TrimSpace(part)
        if part == "" {
            continue
        }
        var item AcceptItem
        params := make(map[string]string)
        segments := strings.Split(part, ";")
        for i, seg := range segments {
            seg = strings.TrimSpace(seg)
            if i == 0 {
                // Type/subtype
                parts := strings.SplitN(seg, "/", 2)
                if len(parts) == 2 {
                    item.Type = strings.TrimSpace(parts[0])
                    item.Subtype = strings.TrimSpace(parts[1])
                } else {
                    // Malformed, skip
                    continue
                }
            } else {
                // Parameter
                kv := strings.SplitN(seg, "=", 2)
                if len(kv) == 2 {
                    key := strings.TrimSpace(kv[0])
                    val := strings.TrimSpace(kv[1])
                    params[key] = val
                }
            }
        }
        if item.Type == "" {
            continue
        }
        item.Params = params
        // Quality
        if qStr, ok := params["q"]; ok {
            if q, err := strconv.ParseFloat(qStr, 64); err == nil {
                item.Quality = q
            }
        } else {
            item.Quality = 1.0
        }
        items = append(items, item)
    }
    // Sort by quality descending
    sortAccept(items)
    return items
}

func sortAccept(items []AcceptItem) {
    for i := 0; i < len(items)-1; i++ {
        for j := i + 1; j < len(items); j++ {
            if items[i].Quality < items[j].Quality {
                items[i], items[j] = items[j], items[i]
            }
        }
    }
}

func ParseAcceptLanguage(header string) []string {
    if header == "" {
        return nil
    }
    var langs []string
    for _, part := range strings.Split(header, ",") {
        part = strings.TrimSpace(part)
        if part == "" {
            continue
        }
        // Strip quality if present
        idx := strings.Index(part, ";")
        if idx != -1 {
            part = part[:idx]
        }
        langs = append(langs, strings.TrimSpace(part))
    }
    return langs
}

func ParseAcceptEncoding(header string) []string {
    if header == "" {
        return nil
    }
    var encodings []string
    for _, part := range strings.Split(header, ",") {
        part = strings.TrimSpace(part)
        if part == "" {
            continue
        }
        idx := strings.Index(part, ";")
        if idx != -1 {
            part = part[:idx]
        }
        encodings = append(encodings, strings.TrimSpace(part))
    }
    return encodings
}
