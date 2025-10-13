## Vocab
* **Enumeration**: The process of systematically listing or counting items, elements, or resources in a specific order. In various contexts, it can refer to different activities, but it generally involves identifying and cataloging information for analysis or management purposes.

##  Frameworks

### OSI (Open Systems Interconnection) Model
A conceptual framework to understand and standardize the functions of a networking system. The layers, from the lowest to the highest, are:

* **1 (Physical Layer)**: Manages the physical connection and transmission of raw bits over a medium.
  * **Data Types**: Raw bits (0s and 1s)
  * **Protocols/Technologies**:
    * Ethernet
    * USB
    * DSL
    * RS-232
* **2 (Data Link Layer)**: Provides node-to-node data transfer and error detection/correction.
  * **Data Types**: Frames
  * **Protocols**:
    * Ethernet (IEEE 802.3)
    * Wi-Fi (IEEE 802.11)
    * PPP (Point-to-Point Protocol)
    * HDLC (High-Level Data Link Control)
* **3 (Network Layer)**: Handles routing of data packets and determines the best path for transmission.
  * **Data Types**: Packets
  * **Protocols**:
    * IP (Internet Protocol)
    * ICMP (Internet Control Message Protocol)
    * IGMP (Internet Group Management Protocol)
* **4 (Transport Layer)**: Ensures reliable data transfer, managing flow control and error recovery.
  * **Data Types**: Segments (TCP) / Datagrams (UDP)
  * **Protocols**:
    * TCP (Transmission Control Protocol)
    * UDP (User Datagram Protocol)
    * SCTP (Stream Control Transmission Protocol)
* **5 (Session Layer)**: Manages sessions and controls the dialog between applications.
  * **Data Types**: Data
  * **Protocols**:
    * NetBIOS
    * RPC (Remote Procedure Call)
    * PPTP (Point-to-Point Tunneling Protocol)
* **6 (Presentation Layer)**: Translates data formats, encrypts data, and ensures readability for applications.
  * **Data Types**: Data (formatted for application use)
  * **Protocols**:
    * JPEG (for images)
    * MPEG (for video)
    * TLS/SSL (for encryption)
* **7 (Application Layer)**: Interfaces directly with end-user applications, providing network services.
  * **Data Types**: Data (application-specific)
  * **Protocols**:
    * HTTP/HTTPS (Hypertext Transfer Protocol)
    * FTP (File Transfer Protocol)
    * SMTP (Simple Mail Transfer Protocol)
    * DNS (Domain Name System)
