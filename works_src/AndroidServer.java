package networkPrivate.android;
import java.io.*;
import java.net.*;
import java.util.*;
class AndroidServer extends AbsServer{
	public static void main(String[] args)throws IOException
	{
		AndroidServer server = new AndroidServer();
		server.execAll(4302, 2, InetAddress.getLocalHost());
	}
	
	//This is run when the thread was created.
	@Override
	public void outerRun(SocketThread t){
		try{
			System.out.println("Client IP Address:" + t.socket.getInetAddress());
		}catch(Exception e){
			System.err.println("ErrorNumber 2:" + e);
			//tryClose(t.socket);
		}
	}
}
