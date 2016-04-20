public abstract class AbsClientE extends AbsClient
{
	//Start client
	@Override
	public void start()
	{
		try{
			sock = new Socket(ip, port);
			sock.setSoTimeout(timeout);
			System.out.println("Connection success");
			new ClientThread().start();
		}catch(Exception e){
			System.err.println("ErrorNumber 1:" + e);
			tryClose(sock);
		}
	}
	
	//This is run when the thread was create
	abstract public void outerRun(ClientThread t);
	/*---For exsample
	{
		try{
			//===Your Code
		}catch(Exception e){
			System.err.println("ErrorNumber 2:" + e);
			tryClose(t.sock);
		}
	}*/
	
	//this is inner class(multi thread)
	class ClientThread extends Thread
	{
		Socket socket;
		
		//initialize
		public ClientThread()
		{
			socket = sock;
		}
		
		//
		public void run()
		{
			outerRun(this);
		}
	}
}