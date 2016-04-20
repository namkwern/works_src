package networkPrivate.android;
import java.io.*;
import java.net.*;
import java.util.*;

public abstract class AbsClient
{
	Socket sock;
	InetAddress ip;
	int port;
	int timeout;
	
	//Execution all
	public void execAll(InetAddress ip, int port, int timeout)throws IOException
	{
		set(ip, port ,timeout);
		System.out.println(toString());
		start();
		process();
	}
	public void execAll(InetAddress ip, int port)throws IOException
	{
		execAll(ip, port ,0);
	}
	
	//socket configration
	public void set(InetAddress ip, int port, int timeout)throws IOException
	{
		this.ip = ip;
		this.port = port;
		this.timeout = timeout;
	}
	public void set(InetAddress ip, int port)throws IOException
	{
		set(ip, port, 0);
	}
	
	//Return configration
	public String toString(){
		String conf = "\n---Client information---\n";
		conf += "Port Number to try to use \t:" + port + "\n";
		conf += "IP Address\t:" + ip + "\n";
		conf += "------------------------";
		return conf;
	}
	
	//Start client
	public void start()
	{
		try{
			sock = new Socket(ip, port);
			sock.setSoTimeout(timeout);
			System.out.println("Connection success");
		}catch(Exception e){
			System.err.println("ErrorNumber 1:" + e);
			tryClose(sock);
		}
	}
	
	//main processing
	abstract public void process();
	/*---For example
		{
			try{
				//===your program
			}catch(Exception e){
				tryClose(sock);
			}
		}
	*/
	
	//---Try to close
	public void tryClose(Socket socket)
	{
		try{
			if(socket != null)socket.close();
		}catch(Exception e){
			System.err.println("ErrorNumber 3" + e);
		}
	}
	
	
	
}
