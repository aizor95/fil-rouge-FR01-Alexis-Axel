# CampusLink 🏫

CampusLink est un projet web front-end dédié à la gestion des infrastructures d'un campus. L'application permet de consulter la disponibilité des salles, l'inventaire des équipements, de suivre l'état des incidents et d'en déclarer de nouveaux.

---

## 📂 Structure du projet

```text
projet-fil-rouge-alexis-axel/
│
├── css/
│   └── style.css          # Feuille de style principale (Responsive Design)
├── img/
│   ├── attention.webp     # Visuel pour incident en cours
│   ├── check.png          # Visuel pour incident terminé
│   └── chrono.png         # Visuel pour incident en cours de réparation
│
├── index.html             # Tableau de bord & Réservation de salles
├── salles.html            # Gestion et état des salles de cours
├── equipements.html       # Inventaire des équipements
├── incidents.html         # Suivi & Formulaire de déclaration d'incidents
└── incident-detail.html   # Fiches détaillées des incidents
```

---

## 🛠️ Technologies utilisées

* **HTML5** : Structure sémantique du contenu (`header`, `main`, `section`, `article`, `footer`).
* **CSS3** : Modern layouts (Flexbox & CSS Grid), design réactif et animations au survol.

---

## 📌 Fonctionnalités des pages

1. **Tableau de bord (`index.html`)**
   * Vue synthétique des statistiques (salles disponibles, incidents actifs).
   * Formulaire de réservation de salle.

2. **Salles (`salles.html`)**
   * Liste des salles avec détails (capacité, étage, équipements intégrés).
   * Indicateurs visuels du statut (`Disponible`, `En maintenance`).

3. **Équipements (`equipements.html`)**
   * Recensement du matériel disponible et repérage des salles associées.

4. **Incidents (`incidents.html`)**
   * Affichage des tickets selon leur état d'avancement.
   * Formulaire complet pour signaler un nouveau problème matériel.

5. **Détails des incidents (`incident-detail.html`)**
   * Description explicative détaillée des pannes signalées.

---

## 👥 Auteurs

Projet réalisé par **Alexis** & **Axel**.