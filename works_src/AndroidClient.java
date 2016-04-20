package networkPrivate.android;
import java.io.*;
import java.net.*;
import java.util.*;

class AndroidClient
{
	public static void main(String[] args)throws IOException
	{
		MakeSocket ms = new MakeSocket();
		ms.execAll(InetAddress.getLocalHost(), 4302);
	}
}
class MakeSocket extends AbsClient{
	public void process()
	{
		try{
			//===your program
		}catch(Exception e){
			tryClose(sock);
		}
	}
}