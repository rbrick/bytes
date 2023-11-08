import { Citizen, Wallet } from "@/types/Wallet";
import Image from "next/image";


export default function WalletComponent({address, tokenBalance, totalStake, pendingRewards, stakedCitizens}: Wallet) {
   
    console.log("wtf", tokenBalance)
    return (
        <div>
            <span>Address: {address}</span><br/>
            <span>Balance: {tokenBalance}</span><br/>
            <span>Total Stake: {totalStake}</span><br/>
            <span>Pending Rewards: {pendingRewards}</span><br/>

            { 
             stakedCitizens &&
              stakedCitizens.map(citizen => {
                console.log("citizen")
                return (
                    <CitizenComponent id={citizen.id} image={citizen.image} season={citizen.season} key={citizen.id}/>
                )
              })
            }
        </div>
    );

}

function CitizenComponent({image, id, season}: Citizen) {
    return (
        <div>
            <Image src={image} width={100} height={100} alt=""/>
            <span>Season: {season}, ID: {id}</span>
        </div>
    )
}