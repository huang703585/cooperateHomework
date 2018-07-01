
public class client
{
    public static void main(String[] args)
    {
	csInterface a = new csInterface();
	System.out.println(a.SignIn("huang","123456"));
	System.out.println(a.SignUp("huang","123456"));
	System.out.println(a.getMaintain());
	System.out.println(a.getContact());
	System.out.println(a.getAccount());
	a.close();

	csInterface b = new csInterface();
	System.out.println(b.SignIn("huang","123456"));
	System.out.println(b.SignUp("huang","123456"));
	System.out.println(b.getMaintain());
	System.out.println(b.getContact());
	System.out.println(b.getAccount());
	b.close();
    }
}
