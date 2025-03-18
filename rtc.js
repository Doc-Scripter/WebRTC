    let localStream;
    let localPeerConnection;
    let remotePeerConnection;

    const localVideo = document.getElementById('localVideo');
    const remoteVideo = document.getElementById('remoteVideo');
    const startButton = document.getElementById('startButton');
    const callButton = document.getElementById('callButton');
    const hangupButton = document.getElementById('hangupButton');

    // Start capturing media
    startButton.addEventListener('click', async () => {
      try {
        localStream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true });
        localVideo.srcObject = localStream;
      } catch (e) {
        console.error('Error accessing media devices.', e);
      }
    });

    // Establish a peer-to-peer call
    callButton.addEventListener('click', async () => {
      // Create peer connections
      localPeerConnection = new RTCPeerConnection();
      remotePeerConnection = new RTCPeerConnection();

      // Exchange ICE candidates between peers
      localPeerConnection.onicecandidate = event => {
        if (event.candidate) {
          remotePeerConnection.addIceCandidate(event.candidate).catch(e => console.error(e));
        }
      };

      remotePeerConnection.onicecandidate = event => {
        if (event.candidate) {
          localPeerConnection.addIceCandidate(event.candidate).catch(e => console.error(e));
        }
      };

      // When remote stream is added, display it in the remote video element
      remotePeerConnection.ontrack = event => {
        if (remoteVideo.srcObject !== event.streams[0]) {
          remoteVideo.srcObject = event.streams[0];
        }
      };

      // Add local stream tracks to the local peer connection
      localStream.getTracks().forEach(track => {
        localPeerConnection.addTrack(track, localStream);
      });

      // Create offer from local peer and set local/remote descriptions
      try {
        const offer = await localPeerConnection.createOffer();
        await localPeerConnection.setLocalDescription(offer);
        await remotePeerConnection.setRemoteDescription(offer);

        const answer = await remotePeerConnection.createAnswer();
        await remotePeerConnection.setLocalDescription(answer);
        await localPeerConnection.setRemoteDescription(answer);
      } catch (err) {
        console.error('Error during offer/answer negotiation:', err);
      }
    });

    // Hang up the call
    hangupButton.addEventListener('click', () => {
      if (localPeerConnection) {
        localPeerConnection.close();
        localPeerConnection = null;
      }
      if (remotePeerConnection) {
        remotePeerConnection.close();
        remotePeerConnection = null;
      }
      console.log('Call ended.');
    });