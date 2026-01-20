import AccountMenu from "./AccountMenu";

const Navbar = () => {

    return (
        <nav className="bg-gray-700 absolute top-0 left-0 right-0 flex justify-between items-center p-4 max-w-full w-full border-b-[0.5px] border-green-400">
            <div className="navbar-left">
            </div>
            <div className="navbar-center">
                <ul className="nav-links list-none flex m-0 p-0">
                    <li className="m-4">
                        <a href="/">Test1</a>
                    </li>
                    <li className="m-4">
                        <a href="/">Test2</a>
                    </li>
                    <li className="m-4">
                        <a href="/">Test3</a>
                    </li>
                    <li className="m-4">
                        <a href="/">Test1</a>
                    </li>
                    <li className="m-4">
                        <a href="/">Test2</a>
                    </li>
                    <li className="m-4">
                        <a href="/">Test3</a>
                    </li>
                    <li className="m-4">
                        <a href="/">Test1</a>
                    </li>
                    <li className="m-4">
                        <a href="/">Test2</a>
                    </li>
                    <li className="m-4">
                        <a href="/">Test3</a>
                    </li>
                </ul>
            </div>
            <div className="flex items-center gap-4">
                <AccountMenu />
            </div>
        </nav>
    );
};

export default Navbar