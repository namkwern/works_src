package com.example.a1503.btserver;

import android.app.Activity;
import android.hardware.Sensor;
import android.hardware.SensorEvent;
import android.hardware.SensorEventListener;
import android.hardware.SensorManager;
import android.os.Bundle;
import android.os.CountDownTimer;
import android.os.Handler;
import android.os.Message;
import android.view.KeyEvent;
import android.view.MenuItem;
import android.view.View;
import android.widget.TextView;
import android.widget.Toast;

import com.example.a1503.btserver.amodule.*;

/**
 * Created by a1503 on 2016/05/25.
 */
public class BTServer extends Activity
{
	TextView text;
	SimpleBluetooth btM;
	SimpleBluetooth btC;
	Activity act = this;
	Debug debug = new Debug(this, true);
	SimpleSensor sensor;
	
	@Override
	public void onCreate(Bundle bundle){
		super.onCreate(bundle);
		LayoutConfigration layout = new LayoutConfigration(this, 30.0f, false);
		layout.getTextView("Bluetooth接続許可");
		
		//final BTCountDown timer = new BTCountDown(300000, 1000);
		View.OnClickListener event = new View.OnClickListener(){
			@Override
			public void onClick(View v){
				bluetoothStart();
			}
		};
		layout.getButton("BT接続", event);
		View.OnClickListener closeevent = new View.OnClickListener(){
			@Override
			public void onClick(View v){
				if(btM != null){
					btM.close();
					btM = null;
				}
				if(btC != null){
					btC.close();
					btC = null;
				}
			}
		};
		layout.getButton("切断", closeevent);
		text = layout.getTextView();
		layout.disp();
		
		sensor = new SimpleSensor(this);
		SensorEventListener sEvent = new SensorEventListener() {
			Ave avey = new Ave();
			Ave avez = new Ave(); 
			@Override
			public void onSensorChanged(SensorEvent event) {
				if(btM != null && btM.isEnabled()) {
					float[] data = event.values.clone();
					avey.add(data[1]);
					avez.add(data[2]);
					if(avey.getCount() == 4) {
						int x = (int)(avey.get() * -40);
						int y = (int)(avez.get() * 40);
						if(x != 0 && y != 0) {
							btM.write(x + "\n");
							btM.write(y + "\n");
						}
					}
				}else{
					if(btM != null && !btM.isEnabled())
						sensor.pause();
				}
				
			}
			@Override
			public void onAccuracyChanged(Sensor sensor, int accuracy) {
				//nothing
			}
		};
		sensor.set(sEvent, Sensor.TYPE_GYROSCOPE, SensorManager.SENSOR_DELAY_FASTEST);
	}
	
	//bluetoothを開始
	void bluetoothStart(){
		boolean flag = false;
		if(btM == null) {
			new BtMThread().start();
			flag = true;
		}
		if(btC == null) {
			new BtCThread().start();
			flag = true;
		}
		if(flag)return;
		
		if(!btM.isEnabled())new BtMThread().start();
		if(!btC.isEnabled())new BtCThread().start();
	}
	
	@Override
	protected void onResume(){
		super.onResume();
		sensor.start();
	}
	@Override
	public void onPause(){
		super.onPause();
		sensor.pause();
	}
	
	
	//音量キー関係
	private boolean volflag = true;
	@Override
	public boolean onKeyDown(int key, KeyEvent e){
		if(key == KeyEvent.KEYCODE_VOLUME_DOWN && btC != null) {
			if(volflag) btC.write("0\n");
			volflag = false;
			return true;
		}
		return super.onKeyDown(key, e);
	}
	@Override
	public boolean onKeyUp(int key, KeyEvent e) {
		if(key == KeyEvent.KEYCODE_VOLUME_DOWN && btC != null) {
			btC.write("1\n");
			volflag = true;
			return true;
		}
		return super.onKeyUp(key, e);
	}
	
	//接続スレッド
	class BtMThread extends Thread
	{
		public void run(){
			if(btM != null){
				btM.close();
				btM = null;
			}
			btM = new SimpleBluetooth(act, "cc896eaa-d8f0-d97a-c432-0301d6921a54", true);
		}
	}
	class BtCThread extends Thread
	{
		public void run(){
			if(btC != null){
				btC.close();
				btC = null;
			}
			btC = new SimpleBluetooth(act, "cc896eaa-d8f0-d97a-c432-0301d6921a55", false);
		}
	} 
	
	/*
	class BTCountDown extends CountDownTimer {
		
		public BTCountDown(long millisInFuture, long countDownInterval) {
			super(millisInFuture, countDownInterval);
		}
		
		@Override
		public void onTick(long millisUntilFinished) {
			String time = "" + millisUntilFinished / 1000;
			countText.setText("終了まであと" + time + "秒");
		}
		
		@Override
		public void onFinish() {
			countText.setText("Bluetooth接続を終了しました。");
		}
	}*/
	
	
}
