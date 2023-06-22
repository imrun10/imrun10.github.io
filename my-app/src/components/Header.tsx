import * as React from "react";
import "@picocss/pico"
import "./components.css";
import Link from "next/link";


export default function Header(){

    return(<nav className="navbar">
        <ul>
            <li>
                  <Link href="/" className="text">
                    <strong>Ututor</strong>
                  </Link>
      
             </li>
         </ul>
      
         <ul>
             <li>
                    <Link href="/" className="text">Book</Link>
             </li>
             <li>
                    <Link href="/pages/about" className="text">About</Link>
             </li>
             <li>
                  <Link href="/pages/contact" className="text">Contact</Link>
             </li>
             <li>
                 <Link href="/pages/contact" className="text"> Sign in </Link>
              </li>
          </ul>
           
      </nav>)
}


