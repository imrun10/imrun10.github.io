import * as React from "react";
import "@picocss/pico"
import "./components.css";
import Link from "next/link";


export default function Header(){

    return(<nav className="navbar">
        <ul>
            <li>
                  <Link href="/">
                    <strong>Ututor</strong>
                  </Link>
      
             </li>
         </ul>
      
         <ul>
             <li>
                    <Link href="/">Book</Link>
             </li>
             <li>
                    <Link href="/pages/about">About</Link>
             </li>
             <li>
                  <Link href="/pages/contact">Contact</Link>
             </li>
             <li>
                  <button id="theme-toggle" type="button"  className="myButton" data-theme-switcher="light">
                    Sign in
                  </button>
              </li>
          </ul>
           
      </nav>)
}


