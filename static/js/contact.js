async function sendContact(ev) {
    ev.preventDefault();

    const senderName = document.getElementById('Name').value;
    const senderEmail = document.getElementById('Email').value;
    const senderPhone = document.getElementById('Phone').value;
    const senderObject = document.getElementById('Object').value;
    const senderMessage = document.getElementById('Message').value;

    const webhookBody = {
      embeds: [{
        title: 'Nouveaux message reçu',
        fields: [
          { name: 'Nom et/ou prénom', value: senderName },
          { name: 'Email', value: senderEmail },
          { name: 'Numero de téléphone', value: senderPhone },
          { name: 'Object', value: senderObject },
          { name: 'Message', value: senderMessage }
        ]
      }],
    };

    const webhookUrl = 'https://discord.com/api/webhooks/1276884415296966721/vZDkliT6Evz9j2mAierFmxYd4z1Fhvx6qWdQ_BwUhMyCp5dfD-JGBMHP7EBBYOCe4MWO';

    const response = await fetch(webhookUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(webhookBody),
    });

    if (response.ok) {
      document.location.href='/contact/recu'
    } else {
      document.location.href='/contact/error'
    }
  }