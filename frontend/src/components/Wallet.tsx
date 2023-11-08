import { Citizen, Wallet } from "@/types/Wallet";


export default function WalletComponent({address, tokenBalance, totalStake, pendingRewards}: Wallet) {
    return (
        <div>
            <span>Address: {address}</span>
            <span>Balance: {tokenBalance}</span>
            <span>Total Stake: {totalStake}</span>
            <span>Pending Rewards: {pendingRewards}</span>
        </div>
    );

}

function CitizenComponent({}: Citizen) {}