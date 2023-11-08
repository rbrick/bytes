"use client";
import Image from "next/image";
import styles from "./page.module.css";
import next from "next";
import { env } from "process";
import Bytes from "@/components/Bytes";
import { useEffect, useState } from "react";
import { BytesPrice } from "@/types/Bytes";
import { Wallet } from "@/types/Wallet";
import WalletComponent from "@/components/Wallet";

function fetchPrice(): Promise<BytesPrice> {
  return new Promise<BytesPrice>((resolve, reject) => {
    fetch(`http://localhost:1337/v1/api/price`, {
      method: "GET",
    }).then((response) => {
      response
        .json()
        .then((jsonObj) => {
          resolve(new BytesPrice(jsonObj.etherPrice, jsonObj.usdPrice));
        })
        .catch((err) => {
          console.log(err);
        });
    });
  });
}

function fetchWallet(address: string): Promise<Wallet> {
  return new Promise<Wallet>((resolve, reject) => {
    fetch(`http://localhost:1337/v1/api/wallet/${address}`, {
      method: "GET",
    }).then((response) => {
      response
        .json()
        .then((jsonObj) => {
          resolve(new Wallet(address, jsonObj));
        })
        .catch((err) => {
          console.log(err);
        });
    });
  });
}

let ignoreRemount = false;

const ADDRESS_PATTERN = /^0x([a-fA-F0-9]){40}$/;

export default function Home() {
  const [bytesPrice, setBytesPrice] = useState(new BytesPrice("0.00", "0.00"));
  const [displayWallet, setDisplayWallet] = useState(false);
  const [walletAddress, setWalletAddress] = useState("");
  const [error, setError] = useState("");
  const [fetchedWallet, setFetchedWallet] = useState({} as Wallet);

  function onInputChange(text: string) {
    setWalletAddress(text);

    setDisplayWallet(false);
  }

  function onClick() {
    if (!ADDRESS_PATTERN.test(walletAddress)) {
      // console.log(walletAddress)
      setError("invalid address");
      return;
    } else {
      setDisplayWallet(true);
      setError("");
    }
  }

  useEffect(() => {
    if (displayWallet) {
        fetchWallet(walletAddress).then((wallet) => {
          setFetchedWallet(wallet);
        });
    }

  }, [displayWallet]);

  useEffect(() => {
    if (!ignoreRemount) {
      fetchPrice().then((price) => {
        setBytesPrice(price);
      });
    }

    let interval = setInterval(() => {
      fetchPrice().then((price) => {
        setBytesPrice(price);
      });
    }, 5 * 60 * 1000);
    // clear interval on unmount
    return () => {
      ignoreRemount = true;
      clearInterval(interval);
    };
  });

  return (
    <main>
      {error.length > 0}
      <div>
        <span style={{ color: "red" }}>{error}</span>
      </div>

      <Bytes
        priceEther={bytesPrice.priceEther}
        priceUSD={bytesPrice.priceUSD}
      />
      <input
        type="text"
        onChange={(e) => onInputChange(e.target.value)}
        id="walletAddress"
      ></input>
      <button onClick={onClick}>View Wallet</button>

      {displayWallet && (<WalletComponent
        address={walletAddress}
        pendingRewards={fetchedWallet.pendingRewards}
        tokenBalance={fetchedWallet.tokenBalance}
        totalStake={fetchedWallet.totalStake}
        stakedCitizens={fetchedWallet.stakedCitizens}
      />)}
    </main>
  );
}
