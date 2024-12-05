ALTER TABLE translation
    DROP CONSTRAINT translation_transcription_id_fkey,
    ADD CONSTRAINT translation_transcription_id_fkey FOREIGN KEY (transcription_id) REFERENCES transcription(id) ON DELETE CASCADE;