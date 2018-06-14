import java.net.*;
import java.io.*;

public class myClient
{
    public static void main(String args[])
    {
	int port = 8080;
	InetAddress address = InetAddress.getLocalHost();
	
	System.out.println(address.getHostName());
    }
}
