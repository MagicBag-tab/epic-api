INSERT INTO sagas (title, description, image_url) VALUES
('The Troy Saga', 'Odysseus ends the Trojan War through cunning and sets the stage for his journey home.', 'https://i.scdn.co/image/ab67616d00001e02b1603c82eefee317a412d36b'),
('The Cyclops Saga', 'Odysseus and his crew encounter Polyphemus and face the consequences of defiance.', 'https://i.scdn.co/image/ab67616d00001e02ffe5e344274eb3ea5abcb77f'),
('The Ocean Saga', 'The voyage begins as the crew battles storms, fate, and divine hostility.', 'https://i.scdn.co/image/ab67616d00001e02aab8e079e15c7ef2f9259814'),
('The Circe Saga', 'The sorceress Circe tests Odysseus as his men are transformed and trust is challenged.', 'https://i.scdn.co/image/ab67616d00001e02f0b8c2fc25572acad60b0b09'),
('The Underworld Saga', 'Odysseus descends into the Underworld to seek truth and confront haunting revelations.', 'https://i.scdn.co/image/ab67616d00001e02fba95a78235cb1d98f516b20'),
('The Thunder Saga', 'The crew defies the gods, leading to devastating punishment and loss.', 'https://i.scdn.co/image/ab67616d00001e0257f349cadd503d500afca227'),
('The Wisdom Saga', 'Lessons, resilience, and emotional bonds shape the path forward.', 'https://i.scdn.co/image/ab67616d00001e02bbb2d972c9892dc91c626a62'),
('The Vengeance Saga', 'Odysseus prepares to reclaim his home and confront those who betrayed him.', 'https://i.scdn.co/image/ab67616d00001e026e20dfbc1ea325ca55a783a5'),
('The Ithaca Saga', 'The final return to Ithaca and the ultimate test of identity, loyalty, and love.', 'https://i.scdn.co/image/ab67616d00001e02eaa0deca83cb62bc2c8b2838');

INSERT INTO songs (saga_id, title, description, singers, image_url) VALUES
-- The Troy Saga (1)
(1, 'The Horse and the Infant', 'The Trojan Horse strategy brings the war to its end.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02b1603c82eefee317a412d36b'),
(1, 'Just a Man', 'Odysseus reflects on his humanity and burdens.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02b1603c82eefee317a412d36b'),
(1, 'Full Speed Ahead', 'The journey home begins with urgency and hope.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02b1603c82eefee317a412d36b'),
(1, 'Open Arms', 'Trust and optimism guide the crew forward.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02b1603c82eefee317a412d36b'),
(1, 'Warrior of the Mind', 'Odysseus embraces intellect as his greatest weapon.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02b1603c82eefee317a412d36b'),

-- The Cyclops Saga (2)
(2, 'Polyphemus', 'The crew encounters the cyclops in his cave.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02ffe5e344274eb3ea5abcb77f'),
(2, 'Survive', 'A desperate struggle to escape and stay alive.', 'Jorge Rivera-Herrans, Full Cast', 'https://i.scdn.co/image/ab67616d00001e02ffe5e344274eb3ea5abcb77f'),
(2, 'Remember Them', 'Odysseus asserts his legacy and identity.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02ffe5e344274eb3ea5abcb77f'),
(2, 'My Goodbye', 'A farewell shaped by consequences and pride.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02ffe5e344274eb3ea5abcb77f'),

-- The Ocean Saga (3)
(3, 'Storm', 'A violent storm threatens the crew at sea.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02aab8e079e15c7ef2f9259814'),
(3, 'Luck Runs Out', 'Confidence fades as danger increases.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02aab8e079e15c7ef2f9259814'),
(3, 'Keep Your Friends Close', 'Trust becomes essential for survival.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02aab8e079e15c7ef2f9259814'),
(3, 'Ruthlessness', 'Mercy is abandoned in the face of power.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02aab8e079e15c7ef2f9259814'),

-- The Circe Saga (4)
(4, 'Puppeteer', 'Circe manipulates and controls the crew.', 'Amber Efe', 'https://i.scdn.co/image/ab67616d00001e02f0b8c2fc25572acad60b0b09'),
(4, 'Wouldn''t You Like', 'Temptation and persuasion take center stage.', 'Amber Efe', 'https://i.scdn.co/image/ab67616d00001e02f0b8c2fc25572acad60b0b09'),
(4, 'Done For', 'A confrontation determines the crew’s fate.', 'Jorge Rivera-Herrans, Amber Efe', 'https://i.scdn.co/image/ab67616d00001e02f0b8c2fc25572acad60b0b09'),
(4, 'There Are Other Ways', 'Compassion offers an alternative path.', 'Amber Efe', 'https://i.scdn.co/image/ab67616d00001e02f0b8c2fc25572acad60b0b09'),

-- The Underworld Saga (5)
(5, 'The Underworld', 'Odysseus enters the realm of the dead.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02fba95a78235cb1d98f516b20'),
(5, 'No Longer You', 'A prophecy reveals a changed future.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02fba95a78235cb1d98f516b20'),
(5, 'Monster', 'Odysseus confronts what he is becoming.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02fba95a78235cb1d98f516b20'),

-- The Thunder Saga (6)
(6, 'Suffering', 'Hunger and desperation take hold.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e0257f349cadd503d500afca227'),
(6, 'Different Beast', 'Odysseus begins to change under pressure.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e0257f349cadd503d500afca227'),
(6, 'Scylla', 'The crew faces a deadly sea monster.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e0257f349cadd503d500afca227'),
(6, 'Mutiny', 'Trust collapses and rebellion rises.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e0257f349cadd503d500afca227'),
(6, 'Thunder Bringer', 'Zeus delivers divine punishment.', 'Luke Holt', 'https://i.scdn.co/image/ab67616d00001e0257f349cadd503d500afca227'),

-- The Wisdom Saga (7)
(7, 'Legendary', 'Stories of Odysseus grow beyond the man.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02bbb2d972c9892dc91c626a62'),
(7, 'Little Wolf', 'Growth and mentorship shape the next generation.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02bbb2d972c9892dc91c626a62'),
(7, 'We''ll Be Fine', 'Hope persists despite uncertainty.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02bbb2d972c9892dc91c626a62'),
(7, 'Love in Paradise', 'Love and illusion intertwine.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02bbb2d972c9892dc91c626a62'),
(7, 'God Games', 'The gods debate Odysseus’ fate.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02bbb2d972c9892dc91c626a62'),

-- The Vengeance Saga (8)
(8, 'Not Sorry for Loving You', 'Love and regret collide.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e026e20dfbc1ea325ca55a783a5'),
(8, 'Dangerous', 'The path forward is filled with risk.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e026e20dfbc1ea325ca55a783a5'),
(8, 'Charybdis', 'Another deadly force threatens survival.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e026e20dfbc1ea325ca55a783a5'),
(8, 'Get in the Water', 'A final confrontation is set in motion.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e026e20dfbc1ea325ca55a783a5'),
(8, 'Six Hundred Strike', 'A decisive and powerful retaliation.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e026e20dfbc1ea325ca55a783a5'),

-- The Ithaca Saga (9)
(9, 'The Challenge', 'Penelope’s test sets the stage for truth.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02eaa0deca83cb62bc2c8b2838'),
(9, 'Hold Them Down', 'The suitors tighten their control.', 'Full Cast', 'https://i.scdn.co/image/ab67616d00001e02eaa0deca83cb62bc2c8b2838'),
(9, 'Odysseus', 'The king reveals himself at last.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02eaa0deca83cb62bc2c8b2838'),
(9, 'I Can''t Help but Wonder', 'Doubt and reflection linger.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02eaa0deca83cb62bc2c8b2838'),
(9, 'Would You Fall in Love with Me Again', 'A final question of love, identity, and reunion.', 'Jorge Rivera-Herrans', 'https://i.scdn.co/image/ab67616d00001e02eaa0deca83cb62bc2c8b2838');