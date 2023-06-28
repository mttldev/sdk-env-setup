package main

import "fmt"

func sdk_version() string {
    return "8.1.1"
}

func sdk_url() string {
    return fmt.Sprintf("https://www.renpy.org/dl/%s/renpy-%s-sdk.zip", sdk_version(), sdk_version())
}

func rapt_url() string {
    return fmt.Sprintf("https://www.renpy.org/dl/%s/renpy-%s-rapt.zip", sdk_version(), sdk_version())
}

func renios_url() string {
    return fmt.Sprintf("https://www.renpy.org/dl/%s/renpy-%s-renios.zip", sdk_version(), sdk_version())
}

func web_url() string {
    return base_web_url()
}
