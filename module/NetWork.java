package module;
import java.net.*;
import java.io.*;
public class NetWork
{
	//Stringをbyte[]のIPアドレスにして返す(失敗したら0.0.0.0)
	public static byte[] toIP(String ip)
	{
		byte[] b = new byte[4];
		try{
			ip = Strop.priRep(ip, "\n", "");
			String[] s = ip.split("\\.");
			for(int n = 0; n < 4; n++)
				b[n] = (byte)((char)Integer.parseInt(s[n]));
			
		}catch(Exception e){
			char c = 0;
			for(int n = 0; n < 4; n++)
			b[n] = (byte)c;
		}finally{
			return b;
		}
	}
	public static InetAddress toInetAddress(String ip) throws IOException
	{
		byte[] b = new byte[4];
		try{
			ip = Strop.priRep(ip, "\n", "");
			String[] s = ip.split("\\.");
			for(int n = 0; n < 4; n++)
				b[n] = (byte)((char)Integer.parseInt(s[n]));
			
		}catch(Exception e){
			char c = 0;
			for(int n = 0; n < 4; n++)
			b[n] = (byte)c;
		}finally{
			return InetAddress.getByAddress(b);
		}
	}
	//クライアントメソッド　port,ip,timeout,request→response
	public static String getResponse(int port, String ipstr,int timeout, String request) throws IOException
	{
		String filestr = "";
		try{
			ipstr = Strop.priRep(ipstr, "\n", "");
			byte[] b = toIP(ipstr);
			Socket cs;
			if(b[0] == b[1] && b[1] == b[2] && b[2] == b[3]){
				InetAddress[] ip = InetAddress.getAllByName(ipstr);
				System.out.println("port:" + port);
				System.out.print("ServerHost/IP:" + ip[0].toString() + "\n");
				System.out.print("\nサーバに接続しています...");
				cs = new Socket(ip[0], port);
			}else{
				InetAddress ip = toInetAddress(ipstr);
				System.out.println("port:" + port);
				System.out.print("ServerIP:" + ip.toString() + "\n");
				System.out.print("\nサーバに接続しています...");
				cs = new Socket(ip, port);
			}
			
			cs.setSoTimeout(timeout);
			
			System.out.println("\r                              ");
			//接続が成功したときの処理
			BufferedReader in = new BufferedReader(new InputStreamReader(cs.getInputStream()));
			PrintWriter out = new PrintWriter(cs.getOutputStream(), true);
			BufferedReader sysin = new BufferedReader(new InputStreamReader(System.in));
			
			out.println(request);
			
			filestr = String.valueOf((char)in.read());
			System.out.println("接続に成功しました\n" + cs.toString() + "\n");
			int i = 0;
			while(true) {
				i = in.read();
				if ( i == -1) break;
					filestr += String.valueOf((char)i);
			}
			
			cs.close();		//閉じる
		}catch(Exception e){
			System.out.print("次の接続エラーが発生しました\n");
			System.out.println(e);
			filestr = "Error";
		}finally{
			System.out.println("接続が切断されました");
			return filestr;
		}
	}
	public static int getPort(String url){
		
		String str = url.substring(0, url.indexOf("://"));
		if(str.equals("https"))return 443;
		if(str.equals("http"))return 80;
		if(str.equals("ftp"))return 20;
		return 0;
	}
	public static String getIP(String url){
		int index = url.indexOf("://") + 3;
		String str = url.substring(index, url.indexOf("/", index));
		return str;
	}
	public static String getOther(String url){
		int index = url.indexOf("://") + 3;
		int index2 = url.indexOf("/", index) + 1;
		String str = url.substring(index2);
		return str;
	}
}