import java.io.*;
import java.net.*;
import java.util.*;
public abstract class AbsServer{
	//Send to innerClass
	int num = 0;
	Socket soc;
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
	
	//Disp configration
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
				soc = serverSocket.accept();		//---Start
				socketList.add(soc);
				System.out.println("Sever get connection from client:" + num);
				new SocketThread().start();			//---"run()"
				num++;
			}
		}catch(Exception e){
			System.err.print("ErrorNumber 1:" + e);
			tryClose(soc);
		}
	}
	
	//---This is run when the thread was created.
	//===Do Override!
	public void outerRun(SocketThread t){
	/*---For exsample
		try{
			//===Your Code
		}catch(Exception e){
			System.err.println("ErrorNumber 2:" + e);
			tryClose(t.socket);
		}
	*/
	}
	
	//---Try to close
	public void tryClose(Socket socket)
	{
		try{
			socket.close();
		}catch(Exception e){
			System.err.println("ErrorNumber 3" + e);
		}
	}
	
	
	
	
	
	//---It's inner class "multi Thread"
	class SocketThread extends Thread{
		int number;
		String name;
		Socket socket;
		//---get AndroidServer.socket
		public SocketThread()
		{
			this.number = num;
			this.name = "" + num;
			this.socket = soc;
		}
		
		//---Make thread by calling "start()"
		public void run()
		{
			outerRun(this);
		}
	}
}