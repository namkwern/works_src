package networkPrivate.android;
import java.io.*;
import java.net.*;
import java.util.*;
public abstract class AbsServer{
	//Send to innerClass
	int num = 0;
	Socket sock;
	//Status
	public ArrayList<Socket> socketList = new ArrayList<Socket>();
	public ServerSocket serverSocket;
	//---Server configration value
	private int port;
	private int limit;
	private InetAddress ip;
	
	
	//Execution all
	public void execAll(int port, int limit, InetAddress ip)throws IOException
	{
		set(port, limit, ip);
		System.out.println(toString());
		start();
	}
	
	//Setting Server
	public void set(int port, int limit, InetAddress ip)throws IOException
	{
		this.port = port;
		this.limit = limit;
		this.ip = ip;
		serverSocket = new ServerSocket(port, limit, ip);
	}
	
	//Return configration
	public String toString(){
		String conf = "\n---Server information---\n";
		conf += "Use Port Number\t:" + port + "\n";
		conf += "IP Address\t:" + ip + "\n";
		conf += "------------------------";
		return conf;
	}
	
	//---Start Server
	public void start()throws IOException
	{
		try{
			System.out.println("Waiting for first connection...");
			while(true){
				sock = serverSocket.accept();		//---Start
				socketList.add(sock);
				new SocketThread().start();			//---"run()"
				num++;
			}
		}catch(Exception e){
			System.err.print("ErrorNumber 1:" + e);
			tryClose(sock);
		}
	}
	
	//---This is run when the thread was created.
	//===Do Override!
	abstract public void outerRun(SocketThread t);
	/*---For exsample
	{
		try{
			//===Your Code
		}catch(Exception e){
			System.err.println("ErrorNumber 2:" + e);
			tryClose(t.socket);
		}
	}*/
	
	//---Try to close
	public void tryClose(Socket socket)
	{
		try{
			if(socket != null)socket.close();
		}catch(Exception e){
			System.err.println("ErrorNumber 3" + e);
		}
	}
	
	
	
	
	
	//---This is inner class "multi Thread"
	class SocketThread extends Thread{
		int number;
		String name;
		Socket socket;
		//---get AndroidServer.socket
		public SocketThread()
		{
			this.number = num;
			this.name = "" + num;
			this.socket = sock;
		}
		
		//---Make thread by calling "start()"
		public void run()
		{
			outerRun(this);
		}
	}
}
interface FixedMethod{
	
}