#!/bin/bash

# Analyze if the script has been run with sudo
function SudoAnalyze 
{
    if [ "$EUID" -ne 0 ]; then 
        echo "Thanks to run the script with sudo."
        TerminateScript
    fi
}

function OptionsAnalyze
{
    # Initialise une variable pour suivre si une option a été traitée
    local optionProcessed=0

    while [ "$1" != "" ]; do
        OPTION=$1
        shift
        case $OPTION in
            -h | --help) optionProcessed=1 && DisplayHelp && TerminateScript;;
            *) echo "Unknown option: $OPTION, please do --help for availables options"; TerminateScript;;
        esac
    done
    # Delete by default
    if [ $optionProcessed -eq 0 ]; then
        Delete
    fi
}

function DisplayHelp
{
    echo "-h | --help ........ To obtain some help about the script"
}


function Delete
{
    # Delete docker compose and volumes for blockchain
    docker compose --profile explorer --profile ether-faucet down
    rm -Rf ./consensus/beacondata ./consensus/validatordata ./consensus/genesis.ssz
    rm -Rf ./execution/geth
}


function TerminateScript
{
    exit 1
}

# Main Actions

SudoAnalyze

OptionsAnalyze "$@"

