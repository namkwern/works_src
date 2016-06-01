package module;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;

public class Command{
	
  	private static InputStream in = null;   	
  	
  	private static InputStream ein = null;
  	
  	private static OutputStream out = null;
  	
  	private static BufferedReader br = null;
  	
  	private static BufferedReader ebr = null;
	
  	private static Process process = null;
	
  	private static String line = null;
	
  	private static String errLine = null;
	
  	private static Thread stdRun  = null;
    
  	private static Thread errRun  = null;    
    

	public Command() {	
	}	
	
	public static void execCmd(String[] cmd) throws IOException, InterruptedException{	 
		try{
			process = Runtime.getRuntime().exec(cmd);
			
			/* 1 サブプロセスの入力ストリームを取得 */
			in = process.getInputStream(); 
			
			/* 2 サブプロセスのエラーストリームを取得 */
			ein = process.getErrorStream();
			
			/* 3 サブプロセスの出力ストリームを取得 ここでは使用しません。*/
			out = process.getOutputStream();
			
			/* 上記の3つのストリームは finally でクローズします。 */
					
			try {
			/*上記 1 のストリームを別スレッドで出力します。*/
			Runnable inputStreamThread = new Runnable(){
				public void run(){		
					try {
						br = 
						new BufferedReader(new InputStreamReader(in));
						while ((line = br.readLine()) != null) {
							System.out.println(line);
						}
					} catch (Exception e) {		
						e.printStackTrace();      	
					}
				}
			};
			/*上記 2 のストリームを別スレッドで出力*/
			Runnable errStreamThread = new Runnable(){
				public void run(){
					try {
						ebr = new BufferedReader(new InputStreamReader(ein));
						while ((errLine = ebr.readLine()) != null) {
							System.err.println(errLine);
						}          	
					} catch (Exception e) {
				  	  e.printStackTrace();
					}          
				}
			};
				
			stdRun = new Thread(inputStreamThread);
			errRun = new Thread(errStreamThread);
				
			/* スレッドを開始します。 */
			stdRun.start();        
			errRun.start();
				
			/*プロセスが終了していなければ終了するまで待機*/
			int c = process.waitFor();
				
			/* サブスレッドが終了するのを待機 */
			stdRun.join();
			errRun.join();
				
			/*プロセス終了コード出力 */
			//System.out.println(c);
				
			} catch (Exception e) {		
			  	e.printStackTrace();		
			}finally{
			  	if(br!=null)br.close();
			  	if(ebr!=null)ebr.close();
				
			  	/* 子プロセス */
			  	if(in!=null)in.close();
			  	if(ein!=null)ein.close();
			  	if(out!=null)out.close();		
			}
		} catch (Exception e) {		
			e.printStackTrace();
		}
	}

}