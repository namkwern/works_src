#include <iostream>
#include <stdio.h>
#include <WinSock2.h>
#include <ws2bth.h>

#pragma comment(lib, "C:\\Program Files (x86)\\Microsoft SDKs\\Windows\\v7.0A\\Lib\\Ws2_32.lib")

//extern "C"{
int hego()
{
	
    WSAData wsaData = { 0 };
    WSAStartup(MAKEWORD(2, 2), &wsaData);

    SOCKET listen_sock = socket(AF_BTH, SOCK_STREAM, BTHPROTO_RFCOMM);
    if (listen_sock == INVALID_SOCKET) {
        return -1;
    }

    SOCKADDR_BTH sa = { 0 };
    sa.addressFamily = AF_BTH;
    sa.port = BT_PORT_ANY;
            
    if (bind(listen_sock, (SOCKADDR *)&sa, sizeof(sa)) == SOCKET_ERROR) {
        return -1;
    }
            
    int size = sizeof(sa);
    getsockname(listen_sock, (SOCKADDR *)&sa, &size);
        
    CSADDR_INFO info = { 0 };
    info.LocalAddr.lpSockaddr = (LPSOCKADDR)&sa;
    info.LocalAddr.iSockaddrLength = sizeof(sa);
    info.iSocketType = SOCK_STREAM;
    info.iProtocol = BTHPROTO_RFCOMM;

    WSAQUERYSET set = { 0 };
    set.dwSize = sizeof(WSAQUERYSET);                              // Must be set to sizeof(WSAQUERYSET)
    set.dwOutputFlags = 0;                                         // Not used
    set.lpszServiceInstanceName = (char*)"Server";                        // Recommended.
    
    GUID guid;//cc896eaa-d8f0-d97a-c432-0301d6921a54
	guid.Data1 = 0xcc896eaa;
	guid.Data2 = 0xd8f0;
	guid.Data3 = 0xd97a;
	guid.Data4[0] = 0xc4;
	guid.Data4[1] = 0x32;
	guid.Data4[2] = 0x03;
	guid.Data4[3] = 0x01;
	guid.Data4[4] = 0xd6;
	guid.Data4[5] = 0x92;
	guid.Data4[6] = 0x1a;
	guid.Data4[7] = 0x54;
	printf("%08x-%04x-%04x", guid.Data1, guid.Data2, guid.Data3);
	int i;
	printf("-");
	for(i = 0; i < sizeof(guid.Data4); i++){
		printf("%02x", guid.Data4[i]);
		if(i == 1)
			printf("-");
	}
	printf("\n");
	
	set.lpServiceClassId = &guid;   // Requred.
	set.lpVersion = NULL;                                          // Not used.
    set.lpszComment = NULL;                                        // Optional.
    set.dwNameSpace = NS_BTH;                                      // Must be NS_BTH.
    set.lpNSProviderId = NULL;                                     // Not required.
    set.lpszContext = NULL;                                        // Not used.
    set.dwNumberOfProtocols = 0;                                   // Not used.
    set.lpafpProtocols = NULL;                                     // Not used.
    set.lpszQueryString = NULL;                                    // not used.
    set.dwNumberOfCsAddrs = 1;                                     // Must be 1.
    set.lpcsaBuffer = &info;                                       // Pointer to a CSADDR_INFO.
    set.lpBlob = NULL;                                             // Optional.

    if (WSASetService(&set, RNRSERVICE_REGISTER, 0) != 0) {
        return -1;
    }
    
    listen(listen_sock, 0);
printf("listenÅ®accept\n");
    SOCKADDR_BTH sab2;
    int ilen = sizeof(sab2);
    SOCKET socket = accept(listen_sock, (SOCKADDR *)&sab2, &ilen);
printf("acceptÅ®recv\n");

    char buf[1024] = { 0 };
	while(0){
	    int res = recv(socket, buf, sizeof(buf), 0);
	    
	    if (res > 0)
	        printf("%s", buf);
		if(res == 0)break;
	}
printf("recvÅ®close\n");
    closesocket(listen_sock);
    closesocket(socket);

    WSACleanup();
    return 0;
}
//}