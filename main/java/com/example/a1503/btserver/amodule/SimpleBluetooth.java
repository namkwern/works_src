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
	BluetoothSocket[] socket;
	InputStream[] in;
	OutputStream[] out;
	InputStream mainIn;
	OutputStream mainOut;
	Debug de = new Debug(true);
	
	//繋ぐ
	public SimpleBluetooth(Activity act){
		this.act = act;
		bt = BluetoothAdapter.getDefaultAdapter();
		if(bt != null) {
			System.out.println("Bluetoothサポート");
			tryPairing();
		}else {
			System.out.println("Bluetooth非サポート");
		}
		tryConnect();
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
				UUID uuid = UUID.fromString("fa87c0d0-afac-11de-8a39-0900200c9a66");
				socket = new BluetoothSocket[devices.length];
				out = new OutputStream[devices.length];
				in = new InputStream[devices.length];
				try {
					for (int n = 0; n < devices.length; n++) {
						socket[n] = devices[n].createRfcommSocketToServiceRecord(uuid);
						out[n] = socket[n].getOutputStream();
						in[n] = socket[n].getInputStream();
						//認証メッセージ
						write(out[n], "Prease respons message!:0001");
					}
						
				} catch (IOException ioex) {
					throw new NullPointerException(getClass() + ":" + ioex);
				}
				String str = "";
				for(BluetoothDevice bDev : devices)
					str += bDev.getName() + ":";
				de.alert(act, "");
			}else{
				de.alert(act, "ペアリングして");
			}
			
		}
	}
	
	//特定のアプリケーションを見る
	private void connectApp(){
		for(int n = 0; n < out.length; n++){
			write(out[n], "Prease respons message!:0001");
					
		}
	}
	
	//入出力ストリームを入手または設定
	public OutputStream getOutputStream(){
		return mainOut;
	}
	public InputStream getInputStream(){
		return mainIn;
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
		if(mainOut != null)
			write(mainOut, str);
	}
	
	public String read(){
		String str = "";
		
		return str;
	}
	/*
	Handler hand = new Handler();
	class InputWait extends Thread{
		@Override
		public void run(){
			read();
		}
	}*/
}
