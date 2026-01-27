import AccountMenu from "./AccountMenu";
import { Label } from "@radix-ui/react-dropdown-menu";
import { Link } from "react-router-dom";

const Navbar = () => {

    return (
        <nav className="bg-linear-to-b from-gray-700 to-gray-500 absolute top-0 left-0 right-0 flex justify-between items-center p-4 max-w-full w-full border-b-2 [border-image:linear-gradient(to_right,var(--color-green-400),var(--color-green-600))_1]">
            <div className="navbar-left">
            </div>
            <div className="navbar-center">
                <ul className="nav-links list-none flex m-0 p-0 gap-20">
                    <li className="m-4">
                        <Link to="/app/about"><Label>About Us</Label></Link>
                    </li>
                    <li className="m-4">
                        <Link to="/app/gallery"><Label>Gallery</Label></Link>
                    </li>
                    <li className="m-4">
                        <Link to="/app/releases"><Label>Releases</Label></Link>
                    </li>
                    <li className="m-4">
                        <Link to="/app/webstore"><Label>Webstore</Label></Link>
                    </li>
                    <li className="m-4">
                        <Link to="/app/news"><Label>News</Label></Link>
                    </li>
                </ul>
            </div>
            <div className="flex items-center gap-4">
                <Label>Logged in as {}</Label>
                <AccountMenu />
            </div>
        </nav>
    );
};

export default Navbar