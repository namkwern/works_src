package networkPrivate.android;
import java.io.*;
import java.net.*;
import java.util.*;
class AndroidServer extends AbsServer{
	public static void main(String[] args)throws IOException
	{
		AndroidServer server = new AndroidServer();
		server.execAll(4203, 2, InetAddress.getLocalHost());
		server.clear();
	}
	
	@Override
	public void afterConnect(){
		
	}
	
	//This is run when the thread was created.
	@Override
	public void outerRun(SocketThread t){
		try{
			System.out.println("Client IP Address:" + t.socket.getInetAddress());
		}catch(Exception e){
			System.err.println(this.getClass() + "ErrorNumber 1:" + e);
			//tryClose(t.socket);
		}
	}
}
