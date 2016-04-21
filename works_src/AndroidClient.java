package networkPrivate.android;
import java.io.*;
import java.net.*;
import java.util.*;

class AndroidClient extends AbsClient
{
	public static void main(String[] args)throws IOException
	{
		AndroidClient client = new AndroidClient();
		client.execAll(InetAddress.getLocalHost(), 4302);
	}
	
	
	@Override
	public void process()
	{
		try{
			//===your program
		}catch(Exception e){
			tryClose(sock);
		}
	}
}
