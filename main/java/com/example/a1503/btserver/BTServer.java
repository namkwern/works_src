package com.example.a1503.btserver;

import android.app.Activity;
import android.os.Bundle;
import android.os.CountDownTimer;
import android.os.Handler;
import android.os.Message;
import android.view.MenuItem;
import android.view.View;
import android.widget.TextView;

import com.example.a1503.btserver.amodule.LayoutConfigration;
import com.example.a1503.btserver.amodule.SimpleBluetooth;

/**
 * Created by a1503 on 2016/05/25.
 */
public class BTServer extends Activity
{
	TextView countText;
	
	@Override
	public void onCreate(Bundle bundle){
		super.onCreate(bundle);
		LayoutConfigration layout = new LayoutConfigration(this, 30.0f);
		layout.getTextView("Bluetooth接続許可");
		
		//final BTCountDown timer = new BTCountDown(300000, 1000);
		View.OnClickListener event = new View.OnClickListener(){
			@Override
			public void onClick(View v){
				bluetoothStart();
				//timer.start();
			}
		};
		layout.getButton("BT", event);
		
		countText = layout.getTextView();
	}
	
	//bluetoothを開始
	void bluetoothStart(){
		SimpleBluetooth bt = new SimpleBluetooth(this);
		//bt.write("ssidwwwひらがな漢字");
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
