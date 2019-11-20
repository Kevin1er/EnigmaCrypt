# ENIGMA

Ce programme permet de décrypter des messages utilisant une version simplifié d'Enigma (sans le tableau de connexion).

---

## Utilisation du programme

**Prérequis** -> Un environement Go fonctionnel.

1. Lancer le programme avec la commande `go run enigma.go` ;
2. Le programme vous demandera de choisir les rotors, ainsi que le réflecteur puis enfin la clé ;
3. Le programme vous demandera finalement quelle message vous voulez tester ;

Le programme va d'abord crypter le message en utilisant la configuration choisie puis essaiera de retrouver cette configuration par force brute.
