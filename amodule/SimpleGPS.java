package com.example.a1503.car_logger.amodule;

import android.Manifest;
import android.app.Activity;
import android.content.Context;
import android.content.pm.PackageManager;
import android.location.LocationListener;
import android.location.LocationManager;
import android.support.v4.app.ActivityCompat;

/**
 * Created by a1503 on 2016/05/09.
 */
public class SimpleGPS {
	LocationListener sEvent;
	LocationManager manager;
	Activity act;
	
	public SimpleGPS(Activity act) {
		this.act = act;
		manager = (LocationManager) act.getSystemService(Context.LOCATION_SERVICE);
	}
	
	public void start(LocationListener sEvent) {
		this.sEvent = sEvent;
		if (permissionCheck(act)) {
			return;
		}
		manager.requestLocationUpdates(LocationManager.NETWORK_PROVIDER, 0, 0, sEvent);
		
	}
	
	public void pause() {
		if (permissionCheck(act)) {
			return;
		}
		manager.removeUpdates(sEvent);
	}
	
	public boolean permissionCheck(Activity act){
		return (ActivityCompat.checkSelfPermission(act, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED && ActivityCompat.checkSelfPermission(act, Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED);
	}
}
