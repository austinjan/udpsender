package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// Log entries
	logs := []string{
		"<5>May 15 15:31:18 jilllocal main: exiting main() ",
		"<5>May 15 14:12:34 localhost main: WARNING: Low disk space on /dev/sda1 ",
		"<5>May 15 14:15:45 pc-1DDEskc main: ERROR: Unable to connect to database server ",
		"<5>May 15 14:18:52 nms1 main: CRITICAL: System overheating detected ",
		"<5>May 15 14:21:17 desktop-e2drE5 main: WARNING: Memory usage exceeds 90% ",
		"<5>May 15 14:25:03 jilllocal main: ERROR: Failed to write to /var/log/system.log ",
		"<5>May 15 14:29:41 host1 main: CRITICAL: Kernel panic - not syncing: Fatal exception ",
		"<5>May 15 14:33:29 root main: WARNING: Network latency above threshold ",
		"<5>May 15 14:37:55 root main: ERROR: Application crashed unexpectedly ",
		"<5>May 15 14:42:18 jilllocal main: CRITICAL: Power supply failure detected ",
		"<5>May 15 14:46:03 jilllocal main: WARNING: Unauthorized access attempt blocked ",
		"<5>May 15 14:10:12 host1 main: INFO: System startup completed successfully",
		"<5>May 15 14:12:24 serverA main: INFO: Scheduled backup finished without errors",
		"<5>May 15 14:15:33 workstation1 main: INFO: User login successful for user 'admin'",
		"<5>May 15 14:18:56 dbserver main: INFO: Network configuration applied successfully",
		"<5>May 15 14:22:14 client1 main: INFO: Application update completed successfully",
		"<5>May 15 14:25:48 router01 main: INFO: File /var/log/syslog rotated successfully",
		"<5>May 15 14:28:22 storage01 main: INFO: Scheduled task 'cleanup' finished without issues",
		"<5>May 15 14:31:39 firewall1 main: INFO: All services restarted successfully",
		"<5>May 15 14:35:47 proxy01 main: INFO: Disk check completed, no errors found",
		"<5>May 15 14:39:12 backupserver main: INFO: Database backup completed successfully",
		"<185>May 15 15:22:49 192.168.127.93:63292 2017-01-19T17:28:09.968Z 192.168.127.44 syslog: Link Status: Port3 link is up, duplex=Full Duplex, speed=1000. ",
		"<5>May 15 15:22:55 root InsertTopo: new topology from: jilllocal",
		"<5>May 15 15:23:05 jilllocal InsertDev: new device: 00-60-E9-2E-BE-E0",
		"<5>May 15 15:23:05 jilllocal InsertDev: new device: 00-60-E9-1A-3B-89",
		"<5>May 15 15:23:05 jilllocal InsertDev: new device: 00-60-E9-20-BE-27",
		"<5>May 15 15:23:05 jilllocal InsertDev: new device: 00-60-E9-01-96-12",
		"<1>May 15 15:23:05 jilllocal InsertDev: new device: 00-60-E9-28-B6-2B",
		"<1>May 15 15:23:05 jilllocal InsertDev: new device: 00-60-E9-26-31-0D",
		"<1>May 15 15:31:22 jilllocal InsertDev: new device: 00-60-E9-20-BE-27",
		"<1>May 15 15:31:23 jilllocal InsertDev: new device: 00-60-E9-01-96-12",
		"<5>May 15 15:40:47 root main: exiting main() root",
		"<1>May 15 15:40:58 root InsertTopo: new topology from: jilllocal",
		"<1>May 15 15:41:43 root InsertDev: new device: 00-60-E9-1A-3B-89",
		"<1>May 15 15:41:43 root InsertDev: new device: 00-60-E9-20-BE-27 ",
		"<134>Jul 17 06:13:37 combo ftpd[23575]: connection from 83.116.207.11 (aml-sfh-3310b.adsl.wanadoo.nl) at Sun Jul 17 06:13:37 2005  ",
		"<134>Jul 17 08:06:12 combo ftpd[23798]: connection from 218.146.61.230 () at Sun Jul 17 08:06:12 2005 ",
		"<134>Jul 17 08:06:14 combo ftpd[23799]: connection from 218.146.61.230 () at Sun Jul 17 08:06:14 2005 ",
		"<134>Jul 17 09:44:07 combo ftpd[23931]: connection from 210.245.165.136 () at Sun Jul 17 09:44:07 2005 ",
		"<134>Jul 17 10:45:07 combo sshd(pam_unix)[24031]: authentication failure; logname= uid=0 euid=0 tty=NODEVssh ruser= rhost=61-220-159-99.hinet-ip.hinet.net  user=root",
		"<134>Jul 17 10:45:07 combo sshd(pam_unix)[24033]: authentication failure; logname= uid=0 euid=0 tty=NODEVssh ruser= rhost=61-220-159-99.hinet-ip.hinet.net  user=root",
		"<134>Jul 17 10:45:07 combo sshd(pam_unix)[24030]: authentication failure; logname= uid=0 euid=0 tty=NODEVssh ruser= rhost=61-220-159-99.hinet-ip.hinet.net  user=root",
		"<134>Jul 17 14:02:49 combo ftpd[24362]: connection from 207.30.238.8 (host8.topspot.net) at Sun Jul 17 14:02:49 2005 ",
		"<134>Jul 17 14:02:49 combo ftpd[24363]: connection from 207.30.238.8 (host8.topspot.net) at Sun Jul 17 14:02:49 2005 "}

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// -logs to send logs boolean
	logsFlag := flag.Bool("logs", false, "send logs")
	// Resolve the UDP address
	serverAddr, err := net.ResolveUDPAddr("udp", ":11113")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Create a UDP connection
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}
	defer conn.Close()

	if *logsFlag {

		for {
			// Select a random log
			log := logs[rand.Intn(len(logs))]
			fmt.Printf("Sending log: %s (%s)\n", log, time.Now().Format("2006-01-02 15:04:05"))
			// Send the log
			_, err := conn.Write([]byte(log + "\n"))
			if err != nil {
				fmt.Println("Error sending log:", err)
				return
			}

			// Sleep for a random duration between 1 and 4 seconds
			time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
		}
	} else {
		// Send message
		message := strings.Join(os.Args[1:], " ")
		fmt.Println("Sending message:", message)
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}
