package main
import (
    "log"
    "net"
    "time"
    "strconv"
)

func main() {
    exec(0)
}

func exec(pos int) {
    var rlen int
    var err error

    remote, err := net.ResolveUDPAddr("udp", "localhost:8888")
    local, err := net.ResolveUDPAddr("udp", ":0")

    if err != nil {
        log.Fatalf("%v\n", err)
    }

    conn, err := net.DialUDP("udp", local, remote)

    if err != nil {
        log.Fatalf("%v\n", err)
    }

    conn.SetDeadline(time.Now().Add(5 * time.Second))

    defer conn.Close()

    s := "user" + strconv.Itoa(pos)

    rlen, err = conn.Write([]byte(s))

    if err != nil {
        log.Printf("Send Error: %v\n", err)
        return
    }

    log.Printf("Send[%d]: %v\n", pos, s)

    buf := make([]byte, 1024)

    rlen, remote, err = conn.ReadFromUDP(buf)

    if err != nil {
        log.Printf("Receive Error: %v\n", err)
        return
    }

    log.Printf("Receive[%d]: %v\n", pos, string(buf[:rlen]), remote)
}