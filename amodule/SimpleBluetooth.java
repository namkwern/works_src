package com.example.a1503.btserver.amodule;

import android.app.Activity;
import android.bluetooth.BluetoothAdapter;
import android.bluetooth.BluetoothDevice;
import android.bluetooth.BluetoothSocket;
import android.content.Intent;
import android.widget.ArrayAdapter;
import android.widget.TextView;
import android.widget.Toast;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.util.Set;
import java.util.UUID;
import java.util.logging.Handler;
import java.util.logging.LogRecord;

/**
 * Created by a1503 on 2016/05/10.
 */
public class SimpleBluetooth {
	BluetoothAdapter bt;
	Activity act;
	final int REQUEST_ENABLE_BLUETOOTH = 1;
	BluetoothSocket socket;
	InputStream in;
	OutputStream out;
	Debug de;
	UUID uuid;
	boolean startFlag = true;
	
	//繋ぐ
	public SimpleBluetooth(Activity act, String uuid){
		this.uuid = UUID.fromString(uuid);
		this.act = act;
		de = new Debug(act, true);
		bt = BluetoothAdapter.getDefaultAdapter();
		if(bt != null) {
			System.out.println("Bluetoothサポート");
			//tryPairing();
			tryConnect();
		}else {
			System.out.println("Bluetooth非サポート");
		}
	}
	
	//時間制限有の接続要求
	private void tryPairing(){
		if(bt.getScanMode() != BluetoothAdapter.SCAN_MODE_CONNECTABLE_DISCOVERABLE){
			Intent intent = new Intent(BluetoothAdapter.ACTION_REQUEST_DISCOVERABLE);
			intent.putExtra(BluetoothAdapter.EXTRA_DISCOVERABLE_DURATION, 300);
			act.startActivity(intent);
		}
	}
	
	//接続されるまで待機
	private void stayPairing(){
		while(!bt.isEnabled());
	}
	
	//Bluetooth接続後、ソケット接続要求
	public void tryConnect(){
		if(bt.isEnabled()) {
			Set<BluetoothDevice> devicelist = bt.getBondedDevices();
			BluetoothDevice[] devices = devicelist.toArray(new BluetoothDevice[0]);
			if(devices.length > 0) {//ペアリング済みなら
				try {
					for (int n = 0; n < devices.length; n++) {
						BluetoothSocket soc = devices[n].createRfcommSocketToServiceRecord(uuid);//"cc896eaa-d8f0-d97a-c432-0301d6921a(54|55)"
						new ConnectThread(soc).start();
					}
				} catch (IOException ioex) {
					throw new NullPointerException(getClass() + ":" + ioex);
				}
				String str = "";
				for(BluetoothDevice bDev : devices)
					str += bDev.getName() + ":";
			}else{
				de.alert("ペアリングデバイス数が0");
			}
			
		}
	}
	
	//入出力ストリームを入手または設定
	public OutputStream getOutputStream(){
		return out;
	}
	public InputStream getInputStream(){
		return in;
	}
	
	//出力
	public void write(OutputStream out, String str){
		char[] ch = str.toCharArray();
		try{
			for(char c: ch)
				out.write((int)c);
		}catch(IOException ioe){
			System.err.println(getClass() + ":" + ioe);
		}
	}
	public void write(String str){
		if(out != null)
			write(out, str);
	}
	//入力
	public String read(){
		String str = "";
		
		return str;
	}
	
	//クラスが有効であるか
	public boolean isEnabled(){
		if(socket != null)
			return socket.isConnected();
		return startFlag;
	}
	//ソケットの切断
	public void close(){
		try {
			if(socket != null)
				socket.close();
		}catch(Exception e){
			
		}
	}
	
	class ConnectThread extends Thread
	{
		BluetoothSocket soc;
		public ConnectThread(BluetoothSocket soc){
			this.soc = soc;
		}
		public void run(){
			try{//接続できるものに接続
				if(!soc.isConnected())
					soc.connect();
				if(soc.isConnected()) {
					startFlag = false;
					socket = soc;
					out = socket.getOutputStream();
					in = socket.getInputStream();
				}
			}catch(Exception e){
				//nothing
			}
		}
	}
}
