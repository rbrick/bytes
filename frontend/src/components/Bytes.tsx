import { use, useEffect, useMemo, useState } from "react";
import process from "process";
import { resolve } from "path";
import { BytesPrice } from "@/types/Bytes";

interface BytesPriceDisplayProps {
  priceEther: string;
  priceUSD: string;

  marketCap: string;
  totalSupply: string;
}

export default function Bytes({
  priceEther,
  priceUSD,
  marketCap,
  totalSupply
}: BytesPriceDisplayProps) {
  return (
    <>
      <div className="card">
        <div className="card-header">Price</div>
        <div className="card-body">
          <h5 className="card-title">
            {priceUSD} / {priceEther} Ξ
          </h5>
        </div>
      </div> 
      
      <br/>

      <div className="card">
            <div className="card-header">
                 Total Supply
            </div>
            <div className="card-body">
                <h5 className="card-title">{totalSupply}</h5>
            </div>
        </div>

      <br/>

      <div className="card">
            <div className="card-header">
                 Market Cap
            </div>
            <div className="card-body">
                <h5 className="card-title">{marketCap}</h5>
            </div>
        </div>
    </>
  );
}
