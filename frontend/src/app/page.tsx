'use client';
import Image from 'next/image'
import styles from './page.module.css'
import next from 'next'
import { env } from 'process'
import Bytes from '@/components/Bytes'
import { useEffect, useState } from 'react';
import { BytesPrice } from '@/types/Bytes';


function fetchPrice(): Promise<BytesPrice> {
  return new Promise<BytesPrice>((resolve, reject) => {
      fetch(`http://localhost:1337/v1/api/price`, {
          method: "GET"
      }).then( (response) => {

          // response.text().then((text) => {
          //     console.log(text)
          // })
          response.json().then( (jsonObj) => {
             resolve(new BytesPrice(jsonObj.etherPrice, jsonObj.usdPrice));
          }).catch( (err) => {
              console.log(err)
          } );
      });
  });

}


let ignoreRemount = false;

export default function Home() {

  const [bytesPrice, setBytesPrice] = useState(new BytesPrice("0.00", "0.00"));

  useEffect(() => {
      if (!ignoreRemount) {
          fetchPrice().then((price) => {
              setBytesPrice(price);
          });
      }

      let interval = setInterval( () => {
          fetchPrice().then((price) => {
              setBytesPrice(price);
          })
      }, 5 * 60 * 1000);
      // clear interval on unmount
      return () => {
          ignoreRemount = true;
          clearInterval(interval);
      };

  });

  return (
    <main>
      <Bytes priceEther={bytesPrice.priceEther} priceUSD={bytesPrice.priceUSD} />
    </main>
  )
}
